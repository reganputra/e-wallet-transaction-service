package api

import (
	"e-wallet-transaction-service/helpers"
	"e-wallet-transaction-service/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	HealthCheckService interfaces.IHealthCheckService
}

func (api *HealthCheck) HealthCheckHandler(c *gin.Context) {
	status, err := api.HealthCheckService.HealthCheckService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": status, "error": err.Error()})
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, status, nil)
}
