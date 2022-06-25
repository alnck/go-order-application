package models

import (
	valueObject "customer-service/src/domain/valueObject"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCustonerRequestModel struct {
	Id      primitive.ObjectID  `json:"id"`
	Name    string              `json:"name"`
	Email   string              `json:"email"`
	Address valueObject.Address `json:"address"`
}
