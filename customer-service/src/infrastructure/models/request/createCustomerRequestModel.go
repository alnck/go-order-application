package models

import (
	valueObject "customer-service/src/domain/valueObject"
)

type CreateCustomerRequestModel struct {
	Name    string              `json:"name"`
	Email   string              `json:"email"`
	Address valueObject.Address `json:"address"`
}
