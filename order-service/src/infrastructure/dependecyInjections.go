package infrastructure

import (
	"order-service/src/infrastructure/interfaces"
	"order-service/src/infrastructure/repository"
	"order-service/src/services"
)

func NewOrderServiceResolve() interfaces.IOrderService {
	return services.NewOrderService(repository.NewOrderRepository())
}
