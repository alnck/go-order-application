package repository

import (
	"context"
	"customer-service/src/domain/entity"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _collectionName = "customers"

type ICustomerRepository interface {
	Create(customer *entity.Customer) error
	Update(customer *entity.Customer) error
	Delete(Id primitive.ObjectID) error
	GetAll(customer interface{}, page int, limit int) error
	GetById(Id primitive.ObjectID, customer interface{}) error
	IsValid(Id primitive.ObjectID) (bool, error)
}

var lock sync.Mutex

type customerRepository struct{}

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
			DBInstance = client.Database("customers")
		}
	}
}

func NewCustomerRepository() ICustomerRepository {
	initDBInstance()

	return &customerRepository{}
}

func (*customerRepository) Create(customer *entity.Customer) error {
	_, err := DBInstance.Collection(_collectionName).
		InsertOne(
			context.Background(),
			&customer,
			options.InsertOne())

	return err
}

func (*customerRepository) Update(customer *entity.Customer) error {
	_, err := DBInstance.Collection(_collectionName).
		UpdateOne(
			context.Background(),
			bson.M{"_id": customer.Id},
			customer)

	return err
}

func (*customerRepository) Delete(Id primitive.ObjectID) error {
	_, err := DBInstance.Collection(_collectionName).DeleteOne(context.Background(), bson.M{"_id": Id})

	return err
}

func (*customerRepository) GetAll(customers interface{}, page int, limit int) error {
	options := new(options.FindOptions)

	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := DBInstance.Collection(_collectionName).Find(context.Background(), bson.M{}, options)

	if err != nil {
		return err
	}

	err = cursor.All(context.Background(), customers)

	return err
}

func (*customerRepository) GetById(Id primitive.ObjectID, customer interface{}) error {

	err := DBInstance.Collection(_collectionName).FindOne(context.Background(), bson.M{"_id": Id}).Decode(customer)
	return err
}

func (*customerRepository) IsValid(Id primitive.ObjectID) (bool, error) {
	count, err := DBInstance.Collection(_collectionName).CountDocuments(context.Background(), bson.M{"_id": Id})
	check := count > 0
	return check, err
}
