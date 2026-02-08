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
		client.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	client.WriteJSON(w, http.StatusOK, "Categories fetched successfully", categories)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	// get category by ID
	category, err := h.service.GetByID(intID)
	if err != nil {
		client.WriteJSON(w, http.StatusNotFound, "Category not found", nil)
		return
	}

	client.WriteJSON(w, http.StatusOK, "Category fetched successfully", category)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// validate request body
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if category.Name == "" {
		client.WriteJSON(w, http.StatusUnprocessableEntity, "Name is required", nil)
		return
	}

	if category.Description == "" {
		client.WriteJSON(w, http.StatusUnprocessableEntity, "Description is required", nil)
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		log.Printf("Error when create category: %s", err.Error())
		client.WriteJSON(w, http.StatusInternalServerError, "Something wrong when create category", nil)
		return
	}

	client.WriteJSON(w, http.StatusCreated, "Category created successfully", category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get path parameter {id}
	id := r.PathValue("id")

	// validate and convert ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if updateCategory.Name == "" {
		client.WriteJSON(w, http.StatusUnprocessableEntity, "Name is required", nil)
		return
	}

	if updateCategory.Description == "" {
		client.WriteJSON(w, http.StatusUnprocessableEntity, "Description is required", nil)
		return
	}

	updateCategory.ID = intID
	err = h.service.Update(&updateCategory)
	if err != nil {
		log.Printf("Error when update category: %s", err.Error())
		client.WriteJSON(w, http.StatusInternalServerError, "Something wrong when update category", nil)
		return
	}

	client.WriteJSON(w, http.StatusOK, "Category updated successfully", updateCategory)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	err = h.service.Delete(intID)
	if err != nil {
		log.Printf("Error when delete category: %s", err.Error())
		client.WriteJSON(w, http.StatusInternalServerError, "Something wrong when delete category", nil)
		return
	}

	client.WriteJSON(w, http.StatusOK, "Category deleted successfully", nil)
}
