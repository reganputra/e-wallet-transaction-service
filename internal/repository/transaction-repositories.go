package repository

import (
	"context"
	"e-wallet-transaction-service/constant"
	"e-wallet-transaction-service/internal/models"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, data *models.Transaction) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *TransactionRepo) GetTransactionByReference(ctx context.Context, reference string, includeRefund bool) (models.Transaction, error) {
	var resp models.Transaction

	sql := r.DB.WithContext(ctx).Where("reference = ?", reference)
	if !includeRefund {
		sql = sql.Where("transaction_type != ?", constant.TransactionTypeRefund)
	}

	err := sql.Last(&resp).Error
	return resp, err
}

func (r *TransactionRepo) UpdateStatusTransaction(ctx context.Context, reference string, status string, additionalInfo string) error {
	return r.DB.WithContext(ctx).Exec("UPDATE transactions SET transaction_status = ?, additional_info = ? WHERE reference = ?", status, additionalInfo, reference).Error
}

func (r *TransactionRepo) GetTransaction(ctx context.Context, userId int) ([]models.Transaction, error) {

	var resp []models.Transaction
	err := r.DB.WithContext(ctx).Where("user_id = ?", userId).Order("id DESC").Find(&resp).Error

	return resp, err

}
