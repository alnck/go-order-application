package interfaces

import (
	"order-service/src/domain/entity"
	request "order-service/src/infrastructure/models/request"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderService interface {
	Create(requestModel request.CreateOrderRequestModel) (interface{}, error)
	Update(requestModel request.UpdateOrderRequestModel) (bool, error)
	Delete(id primitive.ObjectID) (bool, error)
	GetById(id primitive.ObjectID) (entity.Order, error)
	GetAll(page int, limit int) ([]entity.Order, error)
	GetByCustomerId(page int, limit int, customerId primitive.ObjectID) ([]entity.Order, error)
	ChangeStatus(id primitive.ObjectID, status string) (bool, error)
}
