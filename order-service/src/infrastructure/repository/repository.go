package repository

import (
	"context"
	"log"
	"order-service/src/domain/entity"
	"order-service/src/infrastructure/interfaces"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _collectionName = "orders"

var lock sync.Mutex

type orderRepository struct{}

var DBInstance *mongo.Database

func initDBInstance() {
	if DBInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if DBInstance == nil {
			clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") //LINK_MONGODB_DOCKER_URI
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Fatal("⛒ Connection Failed to Database")
				log.Fatal(err)
			}

			// Check the connection
			err = client.Ping(context.TODO(), nil)
			if err != nil {
				log.Fatal("⛒ Connection Failed to Database")
				log.Fatal(err)
			}
			DBInstance = client.Database("orderDB")
		}
	}
}

func NewOrderRepository() interfaces.IOrderRepository {
	initDBInstance()

	return &orderRepository{}
}

func (*orderRepository) Create(order *entity.Order) (interface{}, error) {
	result, err := DBInstance.Collection(_collectionName).
		InsertOne(
			context.Background(),
			&order,
			options.InsertOne())

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (*orderRepository) Update(order *entity.Order) (bool, error) {
	updateDocument := bson.M{
		"$set": bson.M{
			"CustomerId": order.CustomerId,
			"Quantity":   order.Quantity,
			"Price":      order.Price,
			"Name":       order.Name,
			"Email":      order.Email,
			"Address":    order.Address,
			"Product":    order.Product,
			"UpdatedAt":  time.Now().UTC(),
		},
	}

	result, err := DBInstance.Collection(_collectionName).
		UpdateOne(
			context.Background(),
			bson.M{"_id": order.Id},
			updateDocument)

	if err != nil {
		return false, err
	}

	return result.UpsertedCount > 0, nil
}

func (*orderRepository) Delete(Id primitive.ObjectID) (bool, error) {
	result, err := DBInstance.Collection(_collectionName).DeleteOne(context.Background(), bson.M{"_id": Id})

	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}

func (*orderRepository) GetAllByFilter(page int, limit int, filter map[string]interface{}) ([]entity.Order, error) {
	options := new(options.FindOptions)

	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := DBInstance.Collection(_collectionName).Find(context.Background(), filter, options)

	if err != nil {
		return nil, err
	}

	var orders []entity.Order

	err = cursor.All(context.Background(), orders)

	return orders, err
}

func (*orderRepository) GetById(Id primitive.ObjectID) (entity.Order, error) {
	var order entity.Order

	err := DBInstance.Collection(_collectionName).FindOne(context.Background(), bson.M{"_id": Id}).Decode(order)

	return order, err
}

func (*orderRepository) UpdateStatus(id primitive.ObjectID, status string) (bool, error) {
	updateDocument := bson.M{
		"$set": bson.M{
			"Status":    status,
			"UpdatedAt": time.Now().UTC(),
		},
	}

	result, err := DBInstance.Collection(_collectionName).
		UpdateOne(
			context.Background(),
			bson.M{"_id": id},
			updateDocument)

	if err != nil {
		return false, err
	}

	return result.UpsertedCount > 0, nil
}
