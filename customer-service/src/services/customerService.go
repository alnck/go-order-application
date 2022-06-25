package services

import (
	"customer-service/src/domain/entity"
	"customer-service/src/domain/repository"
	request "customer-service/src/infrastructure/models/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ICustomerService interface {
		Create(requestModel request.CreateCustomerRequestModel) error
		Update(requestModel request.UpdateCustonerRequestModel) error
		Delete(id primitive.ObjectID) error
		GetById(id primitive.ObjectID) (entity.Customer, error)
		GetAll(page int, limit int) ([]entity.Customer, error)
		IsValid(id primitive.ObjectID) (bool, error)
	}
	customerService struct {
		Repository repository.ICustomerRepository
	}
)

func NewCustomerService(repository repository.ICustomerRepository) ICustomerService {
	return &customerService{Repository: repository}
}

func (service customerService) Create(requestModel request.CreateCustomerRequestModel) error {
	model := entity.Customer{
		Id:        primitive.NewObjectID(),
		Name:      requestModel.Name,
		Email:     requestModel.Email,
		Address:   requestModel.Address,
		CreatedAt: time.Now().UTC(),
	}
	err := service.Repository.Create(&model)
	return err
}

func (service customerService) Update(requestModel request.UpdateCustonerRequestModel) error {
	model := entity.Customer{
		Id:        requestModel.Id,
		Name:      requestModel.Name,
		Email:     requestModel.Email,
		Address:   requestModel.Address,
		UpdatedAt: time.Now().UTC(),
	}
	err := service.Repository.Update(&model)
	return err
}

func (service customerService) Delete(id primitive.ObjectID) error {
	err := service.Repository.Delete(id)
	return err
}

func (service customerService) GetById(id primitive.ObjectID) (entity.Customer, error) {
	var customer entity.Customer
	err := service.Repository.GetById(id, &customer)
	return customer, err
}

func (service customerService) GetAll(page int, limit int) ([]entity.Customer, error) {
	var customers []entity.Customer
	err := service.Repository.GetAll(&customers, page, limit)
	return customers, err
}

func (service customerService) IsValid(id primitive.ObjectID) (bool, error) {
	isCheck, err := service.Repository.IsValid(id)
	return isCheck, err
}
