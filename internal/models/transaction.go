package models

import "time"

type Transaction struct {
	Id                int
	UserId            int
	Amount            float64 `gorm:"column:amount;type:decimal(15,2)"`
	TransactionType   string  `gorm:"column:transaction_type;type:enum('TOPUP','PURCHASE', 'REFUND')"`
	TransactionStatus string  `gorm:"column:transaction_status;type:enum('PENDING', 'SUCCESS', 'FAILED', 'REVERSED')"`
	ReferenceId       string  `gorm:"column:reference_id;type:varchar(255)"`
	Description       string  `gorm:"column:description;type:varchar(255)"`
	AddtionalInfo     string  `gorm:"column:addtional_info;type:text"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (*Transaction) TableName() string {
	return "transactions"
}
