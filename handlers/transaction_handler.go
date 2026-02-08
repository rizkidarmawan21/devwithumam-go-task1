package handlers

import (
	"codewithumam-go-task1/handlers/dto/request"
	"codewithumam-go-task1/internal/client"
	berr "codewithumam-go-task1/internal/errors"
	"codewithumam-go-task1/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req request.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	transaction, err := h.service.Checkout(req, false)
	if err != nil {
		var bizErr *berr.BusinessError
		if errors.As(err, &bizErr) {
			client.WriteJSON(w, bizErr.HTTPStatus, bizErr.Message, nil)
			return
		}
		log.Printf("Error when create checkout order: %s", err.Error())
		client.WriteJSON(w, http.StatusInternalServerError, "Something went wrong when creating checkout order", nil)
		return
	}

	client.WriteJSON(w, http.StatusCreated, "Checkout created successfully", transaction)
}
