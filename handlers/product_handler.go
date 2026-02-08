package handlers

import (
	"codewithumam-go-task1/handlers/dto/response"
	"codewithumam-go-task1/internal/client"
	"codewithumam-go-task1/models"
	"codewithumam-go-task1/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := h.service.GetAll(name)
	if err != nil {
		client.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	productResponses := response.NewProductResponses(products)
	client.WriteJSON(w, http.StatusOK, "Products fetched successfully", productResponses)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	product, err := h.service.GetByID(intID)
	if err != nil {
		client.WriteJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	productResponse := response.NewProductDetailResponse(product)
	client.WriteJSON(w, http.StatusOK, "Product show successfully", productResponse)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		log.Printf("Error when create product: %s", err.Error())
		client.WriteJSON(w, http.StatusInternalServerError, "Something wrong when create product", nil)
		return
	}

	client.WriteJSON(w, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	product.ID = intID
	err = h.service.Update(&product)
	if err != nil {
		log.Printf("Error when update product: %s", err.Error())
		client.WriteJSON(w, http.StatusInternalServerError, "Something wrong when update product", nil)
		return
	}

	client.WriteJSON(w, http.StatusCreated, "Product update successfully", product)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	err = h.service.Delete(intID)
	if err != nil {
		log.Printf("Error when delete product: %s", err.Error())
		client.WriteJSON(w, http.StatusInternalServerError, "Something wrong when delete product", nil)
		return
	}

	client.WriteJSON(w, http.StatusCreated, "Product delete successfully", nil)
}
