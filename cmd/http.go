package cmd

import (
	"e-wallet-transaction-service/external"
	"e-wallet-transaction-service/helpers"
	"e-wallet-transaction-service/internal/api"
	"e-wallet-transaction-service/internal/interfaces"
	"e-wallet-transaction-service/internal/repository"
	"e-wallet-transaction-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHttp() {

	healthCheckRepo := repository.NewHealthCheckRepo()
	healthCheckSvc := &service.HealthCheck{
		HealthCheckRepository: healthCheckRepo,
	}
	healthCheckApi := api.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	// deps := InitializeDependencies()

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandler)

	// walletV1 := r.Group("/transaction/v1")

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}

type Dependency struct {
	External interfaces.IExternal
}

func InitializeDependencies() Dependency {

	external := &external.External{}

	return Dependency{
		External: external,
	}
}
