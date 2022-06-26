package models

import (
	valueObject "customer-service/src/domain/valueObject"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCustomerRequestModel struct {
	Id      primitive.ObjectID  `json:"id" validate:"required"`
	Name    string              `json:"name" validate:"required"`
	Email   string              `json:"email" validate:"required,email"`
	Address valueObject.Address `json:"address" validate:"required"`
}
