package repositories

import (
	"codewithumam-go-task1/handlers/dto/request"
	"codewithumam-go-task1/internal/errors"
	"codewithumam-go-task1/models"
	"database/sql"
	"fmt"
	"strings"
)

// buildTransactionDetailsBatchInsert builds a single INSERT with multiple VALUES for batch insert.
// Avoids N+1 by using one query instead of one per detail.
func buildTransactionDetailsBatchInsert(transactionID int, details []models.TransactionDetail) (string, []interface{}) {

	// Skip if no item
	if len(details) == 0 {
		return "", nil
	}

	const cols = 5 // transaction_id, product_id, product_name, quantity, subtotal

	valueTemplates := make([]string, 0, len(details))
	args := make([]interface{}, 0, len(details)*cols)

	for i, d := range details {
		base := i * cols
		valueTemplates = append(valueTemplates,
			fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)",
				base+1, base+2, base+3, base+4, base+5),
		)

		args = append(
			args,
			transactionID,
			d.ProductID,
			d.ProductName,
			d.Quantity,
			d.Subtotal,
		)
	}

	query := `
INSERT INTO transaction_details
	(transaction_id, product_id, product_name, quantity, subtotal)
VALUES ` + strings.Join(valueTemplates, ", ")

	return query, args
}

// insertTransactionDetailsLoop inserts details one-by-one (N+1 queries). Kept for reference only.
// Prefer buildTransactionDetailsBatchInsert + single Exec for production.
// func insertTransactionDetailsLoop(tx *sql.Tx, transactionID int, details []models.TransactionDetail) error {
// 	for i := range details {
// 		details[i].TransactionID = transactionID
// 		_, err := tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, product_name, quantity, subtotal) VALUES ($1, $2, $3, $4, $5)",
// 			transactionID, details[i].ProductID, details[i].ProductName, details[i].Quantity, details[i].Subtotal)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(req request.CheckoutRequest) (*models.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range req.Items {
		if item.Quantity < 1 {
			return nil, errors.NewInvalidQuantity(item.ProductID, item.Quantity)
		}

		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, errors.NewProductNotFound(item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if item.Quantity > stock {
			return nil, errors.NewInsufficientStock(item.ProductID, productName, item.Quantity, stock)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})

	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// Batch insert all details in a single query to avoid N+1
	if len(details) > 0 {
		for i := range details {
			details[i].TransactionID = transactionID
		}
		query, args := buildTransactionDetailsBatchInsert(transactionID, details)
		_, err = tx.Exec(query, args...)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
