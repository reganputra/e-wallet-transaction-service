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

	deps := InitializeDependencies()

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandler)

	transactionV1 := r.Group("/transaction/v1")
	transactionV1.POST("/create", deps.MiddlewareValidateToken, deps.Transaction.CreateTransaction)
	transactionV1.GET("/", deps.MiddlewareValidateToken, deps.Transaction.GetTransaction)
	transactionV1.GET("/:reference", deps.MiddlewareValidateToken, deps.Transaction.GetTransactionDetail)
	transactionV1.PUT("/update-status/:reference", deps.MiddlewareValidateToken, deps.Transaction.UpdateStatusTransaction)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}

type Dependency struct {
	Transaction    interfaces.ITransactionHandler
	External       interfaces.IExternal
	TransactionApi interfaces.ITransactionHandler
}

func InitializeDependencies() Dependency {

	external := &external.External{}

	trxRepo := &repository.TransactionRepo{
		DB: helpers.DB,
	}
	trxSvc := &service.TransactionService{
		TransactionRepo: trxRepo,
		External:        external,
	}
	trxHandler := &api.TransactionApi{
		TransactionService: trxSvc,
	}

	return Dependency{
		Transaction: trxHandler,
		External:    external,
	}
}
