package repository

import (
	"context"
	"e-wallet-transaction-service/internal/models"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, data models.Transaction) error {
	return r.DB.WithContext(ctx).Create(&data).Error
}
