package infrastructure

import (
	"customer-service/src/infrastructure/interfaces"
	"customer-service/src/infrastructure/repository"
	"customer-service/src/services"
)

func NewCustomerServiceResolve() interfaces.ICustomerService {
	return services.NewCustomerService(repository.NewCustomerRepository())
}
