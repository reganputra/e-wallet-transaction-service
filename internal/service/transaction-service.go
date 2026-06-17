package service

import (
	"context"
	"e-wallet-transaction-service/constant"
	"e-wallet-transaction-service/helpers"
	"e-wallet-transaction-service/internal/interfaces"
	"e-wallet-transaction-service/internal/models"
	"errors"
)

type TransactionService struct {
	TransactionRepo interfaces.ITransactionRepo
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req models.Transaction) (models.CreateTransactionResponse, error) {

	var resp models.CreateTransactionResponse

	req.TransactionStatus = constant.TransactionStatusPending
	req.Reference = helpers.GenerateReference()

	err := s.TransactionRepo.CreateTransaction(ctx, req)
	if err != nil {
		return resp, errors.New("failed to create transaction")
	}

	resp.Reference = req.Reference
	resp.TransactionStatus = req.TransactionStatus

	return resp, nil
}
