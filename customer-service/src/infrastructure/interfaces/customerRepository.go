package interfaces

import (
	"customer-service/src/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICustomerRepository interface {
	Create(customer *entity.Customer) (interface{}, error)
	Update(customer *entity.Customer) (bool, error)
	Delete(Id primitive.ObjectID) (bool, error)
	GetAll(page int, limit int) ([]entity.Customer, error)
	GetById(Id primitive.ObjectID) (entity.Customer, error)
	IsValid(Id primitive.ObjectID) (bool, error)
}
