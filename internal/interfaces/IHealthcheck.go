package interfaces

type IHealthCheckService interface {
	HealthCheckService() (string, error)
}

type HealthCheckRepository interface {
	HealthCheckRepository() (string, error)
}
