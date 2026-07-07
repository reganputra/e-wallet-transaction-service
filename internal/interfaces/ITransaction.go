package interfaces

import (
	"context"
	"e-wallet-transaction-service/internal/models"

	"github.com/gin-gonic/gin"
)

type ITransactionRepo interface {
	CreateTransaction(ctx context.Context, data *models.Transaction) error
	UpdateStatusTransaction(ctx context.Context, reference string, status string, additionalInfo string) error
	GetTransaction(ctx context.Context, userId int) ([]models.Transaction, error)
	GetTransactionByReference(context.Context, string, bool) (models.Transaction, error)
}

type ITransactionSvc interface {
	CreateTransaction(ctx context.Context, data models.Transaction) (models.CreateTransactionResponse, error)
	UpdateStatusTransaction(ctx context.Context, tokenData models.TokenData, req *models.UpdateStatusTransaction) error
	GetTransaction(ctx context.Context, userId int) ([]models.Transaction, error)
	GetTransactionDetail(ctx context.Context, reference string) (models.Transaction, error)
	RefundTransaction(ctx context.Context, tokenData models.TokenData, req *models.RefundTransaction) (models.CreateTransactionResponse, error)
}

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	UpdateStatusTransaction(c *gin.Context)
	GetTransaction(c *gin.Context)
	GetTransactionDetail(c *gin.Context)
	RefundTransaction(c *gin.Context)
}
