package interfaces

import (
	"customer-service/src/domain/entity"
	request "customer-service/src/infrastructure/models/request"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICustomerService interface {
	Create(requestModel request.CreateCustomerRequestModel) (interface{}, error)
	Update(requestModel request.UpdateCustomerRequestModel) (bool, error)
	Delete(id primitive.ObjectID) (bool, error)
	GetById(id primitive.ObjectID) (entity.Customer, error)
	GetAll(page int, limit int) ([]entity.Customer, error)
	IsValid(id primitive.ObjectID) (bool, error)
}
