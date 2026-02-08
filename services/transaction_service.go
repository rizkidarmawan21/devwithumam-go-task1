package services

import (
	"codewithumam-go-task1/handlers/dto/request"
	"codewithumam-go-task1/models"
	"codewithumam-go-task1/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(req request.CheckoutRequest, useLock bool) (*models.Transaction, error) {
	return s.repo.CreateTransaction(req)
}
