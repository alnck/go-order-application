package entity

import (
	valueobject "order-service/src/domain/valueObject"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id         primitive.ObjectID  `bson:"_id"`
	CustomerId primitive.ObjectID  `bson:"Custtomer_id"`
	Quantity   int                 `bson:"Quantity"`
	Price      float64             `bson:"Price"`
	Status     string              `bson:"Status"`
	Name       string              `bson:"Name"`
	Email      string              `bson:"Email"`
	Address    valueobject.Address `bson:"Address"`
	Product    Product             `bson:"Product"`
	CreatedAt  time.Time           `bson:"CreatedAt"`
	UpdatedAt  time.Time           `bson:"UpdatedAt"`
}
