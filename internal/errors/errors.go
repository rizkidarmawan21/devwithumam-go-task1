package errors

import "fmt"

// BusinessError represents a domain error that should be returned to the client
// with a specific HTTP status and user-friendly message.
type BusinessError struct {
	HTTPStatus int    // HTTP status code (e.g. 404, 422)
	Code       string // Machine-readable code for frontend (e.g. "PRODUCT_NOT_FOUND")
	Message    string // Human-readable message for frontend
}

func (e *BusinessError) Error() string {
	return e.Message
}

// NewProductNotFound returns a BusinessError for missing product.
func NewProductNotFound(productID int) *BusinessError {
	return &BusinessError{
		HTTPStatus: 404,
		Code:       "PRODUCT_NOT_FOUND",
		Message:    fmt.Sprintf("Product with id %d not found", productID),
	}
}

// NewInsufficientStock returns a BusinessError when requested quantity exceeds available stock.
func NewInsufficientStock(productID int, productName string, requested, available int) *BusinessError {
	return &BusinessError{
		HTTPStatus: 422,
		Code:       "INSUFFICIENT_STOCK",
		Message:    fmt.Sprintf("Insufficient stock for product \"%s\" (id %d): requested %d, available %d", productName, productID, requested, available),
	}
}

// NewInvalidQuantity returns a BusinessError when quantity is invalid (e.g. <= 0).
func NewInvalidQuantity(productID int, quantity int) *BusinessError {
	return &BusinessError{
		HTTPStatus: 422,
		Code:       "INVALID_QUANTITY",
		Message:    fmt.Sprintf("Invalid quantity for product id %d: %d (must be at least 1)", productID, quantity),
	}
}
