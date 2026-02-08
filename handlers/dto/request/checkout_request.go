package request

import "codewithumam-go-task1/models"

type CheckoutRequest struct {
	CustomerID   *string               `json:"customer_id"`
	CustomerName *string               `json:"customer_name"`
	TableNumber  *int                  `json:"table_number"`
	Items        []models.CheckoutItem `json:"items"`
}
