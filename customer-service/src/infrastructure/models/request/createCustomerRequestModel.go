package models

import (
	valueObject "customer-service/src/domain/valueObject"
)

type CreateCustonerRequestModel struct {
	Name    string              `json:"name"`
	Email   string              `json:"email"`
	Address valueObject.Address `json:"address"`
}
