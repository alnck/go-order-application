package services

import (
	"order-service/src/domain/entity"
	"order-service/src/infrastructure/interfaces"
	request "order-service/src/infrastructure/models/request"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	Repository interfaces.IOrderRepository
}

func NewOrderService(repository interfaces.IOrderRepository) interfaces.IOrderService {
	return &OrderService{Repository: repository}
}

func (service OrderService) Create(requestModel request.CreateOrderRequestModel) (interface{}, error) {
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

	return service.Repository.Create(&model)
}

func (service OrderService) Update(requestModel request.UpdateOrderRequestModel) (bool, error) {
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

	return service.Repository.Update(&model)
}

func (service OrderService) Delete(id primitive.ObjectID) (bool, error) {
	return service.Repository.Delete(id)
}

func (service OrderService) GetById(id primitive.ObjectID) (entity.Order, error) {
	return service.Repository.GetById(id)
}

func (service OrderService) GetAll(page int, limit int) ([]entity.Order, error) {
	return service.Repository.GetAllByFilter(page, limit, bson.M{})
}

func (service OrderService) GetByCustomerId(page int, limit int, customerId primitive.ObjectID) ([]entity.Order, error) {
	return service.Repository.GetAllByFilter(page, limit, bson.M{"CustomerId": customerId})
}

func (service OrderService) ChangeStatus(id primitive.ObjectID, status string) (bool, error) {
	return service.Repository.UpdateStatus(id, status)
}
