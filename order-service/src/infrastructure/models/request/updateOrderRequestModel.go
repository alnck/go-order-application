package models

import (
	"order-service/src/domain/entity"
	valueobject "order-service/src/domain/valueObject"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateOrderRequestModel struct {
	Id         primitive.ObjectID  `json:"Id" validate:"required"`
	CustomerId primitive.ObjectID  `json:"CusttomerId" validate:"required"`
	Quantity   int                 `json:"Quantity" validate:"required,min=1"`
	Price      float64             `json:"Price" validate:"required,min=1"`
	Name       string              `json:"Name" validate:"required"`
	Email      string              `json:"Email" validate:"required,email"`
	Address    valueobject.Address `json:"Address" validate:"required"`
	Product    entity.Product      `json:"Product" validate:"required"`
}
