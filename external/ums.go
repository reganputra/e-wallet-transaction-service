package external

import (
	"context"
	"e-wallet-transaction-service/constant"
	"e-wallet-transaction-service/external/proto/tokenValidation"
	"e-wallet-transaction-service/internal/models"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

type External struct {
}

func (*External) ValidateToken(ctx context.Context, token string) (models.TokenData, error) {
	var resp models.TokenData

	conn, err := grpc.Dial("localhost:7000", grpc.WithInsecure())
	if err != nil {
		return resp, errors.New("failed to dial ums grpc")
	}
	defer conn.Close()

	client := tokenValidation.NewTokenValidationClient(conn)

	req := &tokenValidation.TokenRequest{
		Token: token,
	}

	response, err := client.ValidateToken(ctx, req)
	if err != nil {
		return resp, errors.New("failed to validate token")
	}
	if response.Message != constant.Success {
		return resp, fmt.Errorf("failed to validate token, message: %s", response.Message)
	}

	resp.UserId = response.Data.UserId
	resp.Username = response.Data.Username
	resp.FullName = response.Data.FullName
	return resp, nil
}
