package models

import "time"

type Transaction struct {
	ID           int                `json:"id"`
	CustomerID   *string            `json:"customer_id"`
	CustomerName *string            `json:"customer_name"`
	TableNumber  *int               `json:"table_number"`
	TotalAmount  int                `json:"total_amount"`
	CreatedAt    time.Time          `json:"created_at"`
	Details      TransactionDetails `json:"details"`
}

type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

type TransactionDetails []TransactionDetail

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
