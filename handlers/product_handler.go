package handlers

import (
	"codewithumam-go-task1/dto"
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
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productResponses := dto.NewProductResponses(products)

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Products fetched successfully", productResponses)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := client.NewResponse(http.StatusBadRequest, "Invalid ID", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	product, err := h.service.GetByID(intID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	productResponse := dto.NewProductDetailResponse(product)

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Product show successfully", productResponse)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := client.NewResponse(http.StatusBadRequest, "Invalid request body", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		// Log error with request ID context using slog
		log.Printf("Error when create product: %s", err.Error())

		response := client.NewResponse(http.StatusInternalServerError, "Something wrong when create product", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := client.NewResponse(http.StatusCreated, "Product created successfully", product)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := client.NewResponse(http.StatusBadRequest, "Invalid ID", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := client.NewResponse(http.StatusBadRequest, "Invalid request body", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	product.ID = intID
	err = h.service.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when update product: %s", err.Error())

		response := client.NewResponse(http.StatusInternalServerError, "Something wrong when update product", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := client.NewResponse(http.StatusCreated, "Product update successfully", product)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := client.NewResponse(http.StatusBadRequest, "Invalid ID", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.service.Delete(intID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when delete product: %s", err.Error())

		response := client.NewResponse(http.StatusInternalServerError, "Something wrong when delete product", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := client.NewResponse(http.StatusCreated, "Product delete successfully", nil)
	json.NewEncoder(w).Encode(response)
}
