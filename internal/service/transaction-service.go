package service

import (
	"context"
	"e-wallet-transaction-service/constant"
	"e-wallet-transaction-service/external"
	"e-wallet-transaction-service/helpers"
	"e-wallet-transaction-service/internal/interfaces"
	"e-wallet-transaction-service/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type TransactionService struct {
	TransactionRepo interfaces.ITransactionRepo
	External        interfaces.IExternal
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req models.Transaction) (models.CreateTransactionResponse, error) {

	var resp models.CreateTransactionResponse

	req.TransactionStatus = constant.TransactionStatusPending
	req.Reference = helpers.GenerateReference()

	err := s.TransactionRepo.CreateTransaction(ctx, &req)
	if err != nil {
		return resp, errors.New("failed to create transaction")
	}

	jsonAdditionalInfor := map[string]interface{}{}
	if req.AdditionalInfo != "" {
		err = json.Unmarshal([]byte(req.AdditionalInfo), &jsonAdditionalInfor)
		if err != nil {
			return resp, errors.New("additional info type invalid")
		}
	}

	resp.Reference = req.Reference
	resp.TransactionStatus = req.TransactionStatus

	return resp, nil
}

func (s *TransactionService) UpdateStatusTransaction(ctx context.Context, tokenData models.TokenData, req *models.UpdateStatusTransaction) error {

	// check transaction by reference
	trx, err := s.TransactionRepo.GetTransactionByReference(ctx, req.Reference, false)
	if err != nil {
		return errors.New("transaction not found")
	}

	// check transaction status flow
	statusValid := false
	mapStatus := constant.MapTransactionStatusFlow[trx.TransactionStatus]
	for i := range mapStatus {
		if mapStatus[i] == req.TransactionStatus {
			statusValid = true
		}
	}
	if !statusValid {
		return fmt.Errorf("transaction flow is not valid, status = %s", req.TransactionStatus)
	}

	//  request update data from wallet service
	reqUpdateBalance := external.UpdateBalance{
		Amount:    trx.Amount,
		Reference: req.Reference,
	}
	if req.TransactionStatus == constant.TransactionStatusReversed {
		reqUpdateBalance.Reference = "REVERSED-" + req.Reference

		now := time.Now()
		expiredReversed := trx.CreatedAt.Add(constant.MaximumReversedDuration)
		if now.After(expiredReversed) {
			return errors.New("reversed transaction duration has expired")
		}
	}

	var errUpdateBalance error

	switch trx.TransactionType {
	case constant.TransactionTypeTopup:
		if req.TransactionStatus == constant.TransactionStatusSuccess {
			_, errUpdateBalance = s.External.CreditBalance(ctx, tokenData.Token, reqUpdateBalance)
		} else if req.TransactionStatus == constant.TransactionStatusReversed {
			_, errUpdateBalance = s.External.DebitBalance(ctx, tokenData.Token, reqUpdateBalance)
		}
	case constant.TransactionTypePurchase:
		if req.TransactionStatus == constant.TransactionStatusSuccess {
			_, errUpdateBalance = s.External.DebitBalance(ctx, tokenData.Token, reqUpdateBalance)
		} else if req.TransactionStatus == constant.TransactionStatusReversed {
			_, errUpdateBalance = s.External.CreditBalance(ctx, tokenData.Token, reqUpdateBalance)
		}
	}

	if errUpdateBalance != nil {
		return fmt.Errorf("failed to update balance wallet: %w", errUpdateBalance)
	}

	// update additional info
	additionalInfo := map[string]interface{}{}
	if trx.AdditionalInfo != "" {
		err = json.Unmarshal([]byte(trx.AdditionalInfo), &additionalInfo)
		if err != nil {
			return fmt.Errorf("failed to parse additional info (%s): %w", trx.AdditionalInfo, err)
		}
	}

	newAdditionalInfo := map[string]interface{}{}
	if req.AdditionalInfo != "" {
		err = json.Unmarshal([]byte(req.AdditionalInfo), &newAdditionalInfo)
		if err != nil {
			return fmt.Errorf("failed to parse new additional info (%s): %w", req.AdditionalInfo, err)
		}
	}

	for key, val := range newAdditionalInfo {
		additionalInfo[key] = val
	}

	byteAdditionalInfoJson, err := json.Marshal(additionalInfo)
	if err != nil {
		return errors.New("failed to marshal merge additional info")
	}

	// update status in db
	err = s.TransactionRepo.UpdateStatusTransaction(ctx, req.Reference, req.TransactionStatus, string(byteAdditionalInfoJson))
	if err != nil {
		return errors.New("failed to update transaction status")
	}

	return nil
}

func (s *TransactionService) GetTransaction(ctx context.Context, userId int) ([]models.Transaction, error) {

	return s.TransactionRepo.GetTransaction(ctx, userId)
}

func (s *TransactionService) GetTransactionDetail(ctx context.Context, reference string) (models.Transaction, error) {

	return s.TransactionRepo.GetTransactionByReference(ctx, reference, true)

}
