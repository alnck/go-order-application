package services

import (
	"customer-service/src/domain/entity"
	request "customer-service/src/infrastructure/models/request"
	response "customer-service/src/infrastructure/models/response"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CreateCustomer(requestModel request.CreateCustomerRequestModel) error {
	model := entity.Customer{
		Id:        uuid.New(),
		Name:      requestModel.Name,
		Email:     requestModel.Email,
		Address:   requestModel.Address,
		CreatedAt: time.Now().UTC(),
	}
	fmt.Print(model)
	//Todo  go repository
	return nil
}

func UpdateCustomer(requestModel request.UpdateCustonerRequestModel) error {
	model := entity.Customer{
		Name:      requestModel.Name,
		Email:     requestModel.Email,
		Address:   requestModel.Address,
		UpdatedAt: time.Now().UTC(),
	}
	fmt.Print(model)
	return nil
}

func DeleteCustomer(id uuid.UUID) error {

	return nil
}

func GetCustomer(id uuid.UUID) (response.CustomerResponseModel, error) {

	return response.CustomerResponseModel{}, nil
}

func GetAllCustomer() ([]response.CustomerResponseModel, error) {

	return []response.CustomerResponseModel{}, nil
}

func ValidationCheckCustomer(id uuid.UUID) error {

	return nil
}
