package entity

import (
	valueObject "customer-service/src/domain/valueObject"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID
	Name      string
	Email     string
	Address   valueObject.Address
	CreatedAt time.Time
	UpdatedAt time.Time
}
