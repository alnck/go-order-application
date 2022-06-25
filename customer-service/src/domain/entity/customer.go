package entity

import (
	valueObject "customer-service/src/domain/valueObject"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID           `bson:"Id"`
	Name      string              `bson:"Name"`
	Email     string              `bson:"Email"`
	Address   valueObject.Address `bson:"Address"`
	CreatedAt time.Time           `bson:"CreatedAt"`
	UpdatedAt time.Time           `bson:"UpdatedAt"`
}
