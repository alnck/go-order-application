package entity

import (
	valueObject "customer-service/src/domain/valueObject"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	Id        primitive.ObjectID  `bson:"_id"`
	Name      string              `bson:"Name"`
	Email     string              `bson:"Email"`
	Address   valueObject.Address `bson:"Address"`
	CreatedAt time.Time           `bson:"CreatedAt"`
	UpdatedAt time.Time           `bson:"UpdatedAt"`
}
