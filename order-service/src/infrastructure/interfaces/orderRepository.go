package interfaces

import (
	"order-service/src/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderRepository interface {
	Create(order *entity.Order) (interface{}, error)
	Update(order *entity.Order) (bool, error)
	Delete(Id primitive.ObjectID) (bool, error)
	GetAllByFilter(page int, limit int, filter map[string]interface{}) ([]entity.Order, error)
	GetById(Id primitive.ObjectID) (entity.Order, error)
	UpdateStatus(id primitive.ObjectID, status string) (bool, error)
}
