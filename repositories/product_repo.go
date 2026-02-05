package repositories

import (
	"codewithumam-go-task1/models"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() (models.Products, error) {
	query := "SELECT id, name, price, stock, category_id, created_at FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // close koneksi ke db

	products := make(models.Products, 0)
	for rows.Next() {
		var p models.Product
		var categoryID sql.NullInt64

		// scan data by interface (category_id bisa NULL)
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			val := int(categoryID.Int64)
			p.CategoryId = &val
		}

		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.price,
			p.stock,
			p.created_at,
			p.category_id,
			c.id,
			c.name,
			c.description,
			c.created_at
		FROM products p
		LEFT JOIN categories c ON c.id = p.category_id
		WHERE p.id = $1
	`

	var p models.Product

	// gunakan tipe nullable untuk kolom category dan category_id (karena LEFT JOIN dan FK bisa NULL)
	var (
		catIDFK        sql.NullInt64
		catID          sql.NullInt64
		catName        sql.NullString
		catDescription sql.NullString
		catCreatedAt   sql.NullTime
	)

	err := repo.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.Stock,
		&p.CreatedAt,
		&catIDFK,
		&catID,
		&catName,
		&catDescription,
		&catCreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	// set CategoryId jika FK-nya tidak NULL
	if catIDFK.Valid {
		val := int(catIDFK.Int64)
		p.CategoryId = &val
	}

	// jika category ada (tidak NULL), map ke struct Category
	if catID.Valid {
		p.Category = &models.Category{
			ID:          int(catID.Int64),
			Name:        catName.String,
			Description: catDescription.String,
			CreatedAt:   catCreatedAt.Time,
		}
	}

	return &p, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID, &product.CreatedAt)
	return err
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
