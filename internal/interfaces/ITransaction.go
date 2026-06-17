package interfaces

import (
	"context"
	"e-wallet-transaction-service/internal/models"

	"github.com/gin-gonic/gin"
)

type ITransactionRepo interface {
	CreateTransaction(ctx context.Context, data models.Transaction) error
}

type ITransactionSvc interface {
	CreateTransaction(ctx context.Context, data models.Transaction) (models.CreateTransactionResponse, error)
}

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
}
