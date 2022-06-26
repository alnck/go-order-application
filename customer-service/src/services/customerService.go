package services

import (
	"customer-service/src/domain/entity"
	"customer-service/src/infrastructure/interfaces"
	request "customer-service/src/infrastructure/models/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type customerService struct {
	Repository interfaces.ICustomerRepository
}

func NewCustomerService(repository interfaces.ICustomerRepository) interfaces.ICustomerService {
	return &customerService{Repository: repository}
}

func (service customerService) Create(requestModel request.CreateCustomerRequestModel) (interface{}, error) {
	model := entity.Customer{
		Id:        primitive.NewObjectID(),
		Name:      requestModel.Name,
		Email:     requestModel.Email,
		Address:   requestModel.Address,
		CreatedAt: time.Now().UTC(),
	}

	return service.Repository.Create(&model)
}

func (service customerService) Update(requestModel request.UpdateCustomerRequestModel) (bool, error) {
	model := entity.Customer{
		Id:        requestModel.Id,
		Name:      requestModel.Name,
		Email:     requestModel.Email,
		Address:   requestModel.Address,
		UpdatedAt: time.Now().UTC(),
	}

	return service.Repository.Update(&model)
}

func (service customerService) Delete(id primitive.ObjectID) (bool, error) {
	return service.Repository.Delete(id)
}

func (service customerService) GetById(id primitive.ObjectID) (entity.Customer, error) {
	return service.Repository.GetById(id)
}

func (service customerService) GetAll(page int, limit int) ([]entity.Customer, error) {
	return service.Repository.GetAll(page, limit)
}

func (service customerService) IsValid(id primitive.ObjectID) (bool, error) {
	return service.Repository.IsValid(id)
}
