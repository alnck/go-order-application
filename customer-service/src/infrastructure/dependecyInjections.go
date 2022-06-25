package infrastructure

import (
	"customer-service/src/infrastructure/repository"
	"customer-service/src/services"
)

func NewCustomerServiceResolve() services.ICustomerService {
	return services.NewCustomerService(repository.NewCustomerRepository())
}
