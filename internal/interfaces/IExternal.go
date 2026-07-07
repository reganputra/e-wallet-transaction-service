package interfaces

import (
	"context"
	"e-wallet-transaction-service/external"
	"e-wallet-transaction-service/internal/models"
)

type IExternal interface {
	ValidateToken(ctx context.Context, token string) (models.TokenData, error)
	CreditBalance(ctx context.Context, token string, req external.UpdateBalance) (*external.UpdateBalanceResponse, error)
	DebitBalance(ctx context.Context, token string, req external.UpdateBalance) (*external.UpdateBalanceResponse, error)
}
