package category

import (
	"codewithumam-go-task1/internal/client"
	"encoding/json"
	"net/http"
	"strconv"
)

type httpHandler struct{}

func NewHTTPHandler() *httpHandler {
	return &httpHandler{}
}

func (h *httpHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	categories := CategoriesData

	response := client.NewResponse(http.StatusOK, "Categories fetched successfully", categories)
	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
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
	category := CategoriesData.GetByID(intID)
	if category == nil {
		w.WriteHeader(http.StatusNotFound)
		response := client.NewResponse(http.StatusNotFound, "Category not found", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Category fetched successfully", category)
	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	// validate request body
	var category Category
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
	newCategory := Category{
		ID:          len(CategoriesData) + 1,
		Name:        category.Name,
		Description: category.Description,
	}

	// this method add to memory
	CategoriesData = append(CategoriesData, &newCategory)

	w.WriteHeader(http.StatusCreated)
	response := client.NewResponse(http.StatusCreated, "Category created successfully", newCategory)
	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
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
	category := CategoriesData.GetByID(intID)
	if category == nil {
		w.WriteHeader(http.StatusNotFound)
		response := client.NewResponse(http.StatusNotFound, "Category not found", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// validate request body
	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
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

	// update category
	category.Name = updateCategory.Name
	category.Description = updateCategory.Description

	// this method update to memory
	CategoriesData[intID-1] = category

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Category updated successfully", category)
	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
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
	category := CategoriesData.GetByID(intID)
	if category == nil {
		w.WriteHeader(http.StatusNotFound)
		response := client.NewResponse(http.StatusNotFound, "Category not found", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	// delete category
	// get all element before and after the category
	// and then append to the new CategoriesData
	CategoriesData = append(CategoriesData[:intID-1], CategoriesData[intID:]...)

	w.WriteHeader(http.StatusOK)
	response := client.NewResponse(http.StatusOK, "Category deleted successfully", nil)
	json.NewEncoder(w).Encode(response)
}
