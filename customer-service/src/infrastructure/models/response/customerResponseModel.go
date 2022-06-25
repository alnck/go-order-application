package model

import (
	valueObject "customer-service/src/domain/valueObject"
	"time"

	"github.com/google/uuid"
)

type CustomerResponseModel struct {
	Id        uuid.UUID           `json:"id"`
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	Address   valueObject.Address `json:"address"`
	CreatedAt time.Time           `json:"created_date"`
	UpdatedAt time.Time           `json:"update_date"`
}
