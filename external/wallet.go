package external

import (
	"bytes"
	"context"
	"e-wallet-transaction-service/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type UpdateBalance struct {
	Reference string  `json:"reference"`
	Amount    float64 `json:"amount"`
}

type UpdateBalanceResponse struct {
	Message string `json:"message"`
	Data    struct {
		Balance float64 `json:"balance"`
	} `json:"data"`
}

func (*External) CreditBalance(ctx context.Context, token string, req UpdateBalance) (*UpdateBalanceResponse, error) {

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("failed marshal payload")
	}

	url := helpers.GetEnv("WALLET_HOST", "") + helpers.GetEnv("WALLET_ENDPOINT_CREDIT", "")
	httpReq, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.New("failed create http request")
	}
	httpReq.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.New("failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to credit wallet, status code: %d", resp.StatusCode)
	}

	result := &UpdateBalanceResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors.New("failed to decode response")
	}
	defer resp.Body.Close()
	return result, nil
}

func (*External) DebitBalance(ctx context.Context, token string, req UpdateBalance) (*UpdateBalanceResponse, error) {

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("failed marshal payload")
	}

	url := helpers.GetEnv("WALLET_HOST", "") + helpers.GetEnv("WALLET_ENDPOINT_DEBIT", "")
	httpReq, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.New("failed create http request")
	}
	httpReq.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.New("failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to debit wallet, status code: %d", resp.StatusCode)
	}

	result := &UpdateBalanceResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors.New("failed to decode response")
	}
	defer resp.Body.Close()
	return result, nil
}
