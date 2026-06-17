package api

import (
	"e-wallet-transaction-service/constant"
	"e-wallet-transaction-service/helpers"
	"e-wallet-transaction-service/internal/interfaces"
	"e-wallet-transaction-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionApi struct {
	TransactionService interfaces.ITransactionSvc
}

func (h *TransactionApi) CreateTransaction(c *gin.Context) {

	var (
		log = helpers.Logger
		req models.Transaction
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("Token claim not found in context")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Token claim not found", nil)
	}
	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("Token claim is not of type models.TokenData")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Invalid token claim", nil)
		return
	}

	if !constant.MapTransaction[req.TransactionType] {
		log.Error("Invalid transaction type")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
		return
	}

	req.UserId = int(tokenData.UserId)

	resp, err := h.TransactionService.CreateTransaction(c.Request.Context(), req)
	if err != nil {
		log.Error("failed to create transaction", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constant.ErrFailedBadRequest, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constant.Success, resp)
}
