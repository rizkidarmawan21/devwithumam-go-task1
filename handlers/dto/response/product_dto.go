package response

import "codewithumam-go-task1/models"

type ProductResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	CreatedAt string `json:"created_at"`
}

type ProductDetailResponse struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	Price     int              `json:"price"`
	Stock     int              `json:"stock"`
	CreatedAt string           `json:"created_at"`
	Category  *models.Category `json:"category"`
}

// NewProductResponses mengubah list Product model ke list ProductResponse DTO
func NewProductResponses(products models.Products) []ProductResponse {
	responses := make([]ProductResponse, 0, len(products))
	for _, p := range products {
		responses = append(responses, ProductResponse{
			ID:        p.ID,
			Name:      p.Name,
			Price:     p.Price,
			Stock:     p.Stock,
			CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return responses
}

// NewProductDetailResponse mengubah Product model ke ProductDetailResponse DTO
func NewProductDetailResponse(p *models.Product) ProductDetailResponse {
	return ProductDetailResponse{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Stock:     p.Stock,
		CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		Category:  p.Category,
	}
}

