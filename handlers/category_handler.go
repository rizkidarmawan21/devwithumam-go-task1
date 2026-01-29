package handlers

import (
	"codewithumam-go-task1/internal/client"
	"codewithumam-go-task1/models"
	"codewithumam-go-task1/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Categories fetched successfully", categories)
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
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

	// get category by ID
	category, err := h.service.GetByID(intID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := client.NewResponse(http.StatusNotFound, "Category not found", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Category fetched successfully", category)
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// validate request body
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := client.NewResponse(http.StatusBadRequest, "Invalid request body", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// validate request body
	if category.Name == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := client.NewResponse(http.StatusUnprocessableEntity, "Name is required", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// validate request body
	if category.Description == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := client.NewResponse(http.StatusUnprocessableEntity, "Description is required", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// create category
	err = h.service.Create(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when create category: %s", err.Error())
		response := client.NewResponse(http.StatusInternalServerError, "Something wrong when create category", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := client.NewResponse(http.StatusCreated, "Category created successfully", category)
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
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

	// validate request body
	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := client.NewResponse(http.StatusBadRequest, "Invalid request body", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// validate request body
	if updateCategory.Name == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := client.NewResponse(http.StatusUnprocessableEntity, "Name is required", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// validate request body
	if updateCategory.Description == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response := client.NewResponse(http.StatusUnprocessableEntity, "Description is required", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// update category
	updateCategory.ID = intID
	err = h.service.Update(&updateCategory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when update category: %s", err.Error())
		response := client.NewResponse(http.StatusInternalServerError, "Something wrong when update category", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Category updated successfully", updateCategory)
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
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

	// delete category
	err = h.service.Delete(intID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when delete category: %s", err.Error())
		response := client.NewResponse(http.StatusInternalServerError, "Something wrong when delete category", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Category deleted successfully", nil)
	json.NewEncoder(w).Encode(response)
}
