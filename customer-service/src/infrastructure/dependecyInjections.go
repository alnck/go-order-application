package infrastructure

import (
	"customer-service/src/domain/repository"
	"customer-service/src/services"
)

func NewCustomerServiceResolve() services.ICustomerService {
	return services.NewCustomerService(repository.NewCustomerRepository())
}
