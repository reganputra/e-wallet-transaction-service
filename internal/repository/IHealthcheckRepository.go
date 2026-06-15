package repository

type HealthCheckRepo struct{}

func NewHealthCheckRepo() *HealthCheckRepo {
	return &HealthCheckRepo{}
}

func (r *HealthCheckRepo) HealthCheckRepository() (string, error) {
	return "OK", nil
}
