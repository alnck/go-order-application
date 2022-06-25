package services

import (
	"order-service/src/domain/entity"
	request "order-service/src/infrastructure/models/request"
	"order-service/src/infrastructure/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IOrderService interface {
		Create(requestModel request.CreateOrderRequestModel) error
		Update(requestModel request.UpdateOrderRequestModel) error
		Delete(id primitive.ObjectID) error
		GetById(id primitive.ObjectID) (entity.Order, error)
		GetAll(page int, limit int) ([]entity.Order, error)
		ChangeStatus(id primitive.ObjectID, status string) error
	}
	OrderService struct {
		Repository repository.IOrderRepository
	}
)

func NewOrderService(repository repository.IOrderRepository) IOrderService {
	return &OrderService{Repository: repository}
}

func (service OrderService) Create(requestModel request.CreateOrderRequestModel) error {
	model := entity.Order{
		Id:         primitive.NewObjectID(),
		CustomerId: requestModel.CustomerId,
		Quantity:   requestModel.Quantity,
		Price:      requestModel.Price,
		Status:     "New",
		Name:       requestModel.Name,
		Email:      requestModel.Email,
		Address:    requestModel.Address,
		Product: entity.Product{
			Id:       primitive.NewObjectID(),
			ImageUrl: requestModel.Product.ImageUrl,
			Name:     requestModel.Product.Name,
		},
		UpdatedAt: time.Now().UTC(),
	}
	err := service.Repository.Create(&model)
	return err
}

func (service OrderService) Update(requestModel request.UpdateOrderRequestModel) error {
	model := entity.Order{
		Id:         requestModel.Id,
		CustomerId: requestModel.CustomerId,
		Quantity:   requestModel.Quantity,
		Price:      requestModel.Price,
		Name:       requestModel.Name,
		Email:      requestModel.Email,
		Address:    requestModel.Address,
		Product:    requestModel.Product,
		UpdatedAt:  time.Now().UTC(),
	}
	err := service.Repository.Update(&model)
	return err
}

func (service OrderService) Delete(id primitive.ObjectID) error {
	err := service.Repository.Delete(id)
	return err
}

func (service OrderService) GetById(id primitive.ObjectID) (entity.Order, error) {
	var order entity.Order
	err := service.Repository.GetById(id, &order)
	return order, err
}

func (service OrderService) GetAll(page int, limit int) ([]entity.Order, error) {
	var orders []entity.Order
	err := service.Repository.GetAll(&orders, page, limit)
	return orders, err
}

func (service OrderService) ChangeStatus(id primitive.ObjectID, status string) error {
	model := entity.Order{
		Id:        id,
		Status:    status,
		UpdatedAt: time.Now().UTC(),
	}
	err := service.Repository.Update(&model)
	return err
}
