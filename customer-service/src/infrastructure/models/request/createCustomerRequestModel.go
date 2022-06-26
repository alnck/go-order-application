package models

import (
	valueObject "customer-service/src/domain/valueObject"
)

type CreateCustomerRequestModel struct {
	Name    string              `json:"name" validate:"required"`
	Email   string              `json:"email" validate:"required,email"`
	Address valueObject.Address `json:"address" validate:"required"`
}
