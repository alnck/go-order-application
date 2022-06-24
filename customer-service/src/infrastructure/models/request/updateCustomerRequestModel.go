package models

import (
	valueObject "customer-service/src/domain/valueObject"

	"github.com/google/uuid"
)

type UpdateCustonerRequestModel struct {
	Id      uuid.UUID           `json:"id"`
	Name    string              `json:"name"`
	Email   string              `json:"email"`
	Address valueObject.Address `json:"address"`
}
