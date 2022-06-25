package repository

import (
	"context"
	"log"
	"order-service/src/domain/entity"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _collectionName = "orders"

type IOrderRepository interface {
	Create(order *entity.Order) error
	Update(order *entity.Order) error
	Delete(Id primitive.ObjectID) error
	GetAll(order interface{}, page int, limit int) error
	GetById(Id primitive.ObjectID, order interface{}) error
}

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
			DBInstance = client.Database("orders")
		}
	}
}

func NewOrderRepository() IOrderRepository {
	initDBInstance()

	return &orderRepository{}
}

func (*orderRepository) Create(order *entity.Order) error {
	_, err := DBInstance.Collection(_collectionName).
		InsertOne(
			context.Background(),
			&order,
			options.InsertOne())

	return err
}

func (*orderRepository) Update(order *entity.Order) error {

	updateDocument := bson.M{
		"$set": bson.M{
			"Name":      order.Name,
			"Email":     order.Email,
			"Address":   order.Address,
			"UpdatedAt": time.Now().UTC(),
		},
	}
	_, err := DBInstance.Collection(_collectionName).
		UpdateOne(
			context.Background(),
			bson.M{"_id": order.Id},
			updateDocument)

	return err
}

func (*orderRepository) Delete(Id primitive.ObjectID) error {
	_, err := DBInstance.Collection(_collectionName).DeleteOne(context.Background(), bson.M{"_id": Id})

	return err
}

func (*orderRepository) GetAll(orders interface{}, page int, limit int) error {
	options := new(options.FindOptions)

	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := DBInstance.Collection(_collectionName).Find(context.Background(), bson.M{}, options)

	if err != nil {
		return err
	}

	err = cursor.All(context.Background(), orders)

	return err
}

func (*orderRepository) GetById(Id primitive.ObjectID, order interface{}) error {

	err := DBInstance.Collection(_collectionName).FindOne(context.Background(), bson.M{"_id": Id}).Decode(order)
	return err
}
