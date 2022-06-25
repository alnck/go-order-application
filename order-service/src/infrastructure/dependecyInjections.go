package infrastructure

import (
	"order-service/src/infrastructure/repository"
	"order-service/src/services"
)

func NewOrderServiceResolve() services.IOrderService {
	return services.NewOrderService(repository.NewOrderRepository())
}
