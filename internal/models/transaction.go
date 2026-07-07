package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	Id                int       `json:"id"`
	UserId            int       `json:"user_id" validate:"required"`
	Amount            float64   `json:"amount" gorm:"column:amount;type:decimal(15,2)" validate:"required"`
	TransactionType   string    `json:"transaction_type" gorm:"column:transaction_type;type:enum('TOPUP','PURCHASE', 'REFUND')" validate:"required"`
	TransactionStatus string    `json:"transaction_status" gorm:"column:transaction_status;type:enum('PENDING', 'SUCCESS', 'FAILED', 'REVERSED')" validate:"required"`
	Reference         string    `json:"reference" gorm:"column:reference;type:varchar(255)"`
	Description       string    `json:"description" gorm:"column:description;type:varchar(255)"`
	AdditionalInfo    string    `json:"additional_info" gorm:"column:additional_info;type:text"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}

func (l Transaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

func (*Transaction) TableName() string {
	return "transactions"
}

type CreateTransactionResponse struct {
	Reference         string `json:"reference"`
	TransactionStatus string `json:"transaction_status"`
}

type UpdateStatusTransaction struct {
	Reference         string `json:"reference" validate:"required"`
	TransactionStatus string `json:"transaction_status" validate:"required"`
	AdditionalInfo    string `json:"additional_info"`
}

func (l UpdateStatusTransaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
