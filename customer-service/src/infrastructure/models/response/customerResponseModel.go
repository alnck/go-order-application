package model

import (
	valueObject "customer-service/src/domain/valueObject"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerResponseModel struct {
	Id        primitive.ObjectID  `json:"id"`
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	Address   valueObject.Address `json:"address"`
	CreatedAt time.Time           `json:"created_date"`
	UpdatedAt time.Time           `json:"update_date"`
}
