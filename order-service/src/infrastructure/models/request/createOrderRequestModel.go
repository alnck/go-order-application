package models

import (
	"order-service/src/domain/entity"
	valueobject "order-service/src/domain/valueObject"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderRequestModel struct {
	CustomerId primitive.ObjectID  `json:"CusttomerId"`
	Quantity   int                 `json:"Quantity"`
	Price      float64             `json:"Price"`
	Name       string              `json:"Name"`
	Email      string              `json:"Email"`
	Address    valueobject.Address `json:"Address"`
	Product    entity.Product      `json:"Product"`
}
