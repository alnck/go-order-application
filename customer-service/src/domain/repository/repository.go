package repository

import (
	"context"
	"customer-service/src/domain/entity"
	"log"
	"sync"

	response "customer-service/src/infrastructure/models/response"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _collectionName = "customers"

type ICustomerRepository interface {
	Create(customer *entity.Customer) error
	Update(customer *entity.Customer) error
	Delete(Id uuid.UUID) error
	GetAll() (*[]response.CustomerResponseModel, error)
	GetById(Id uuid.UUID) error
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
			DBInstance = client.Database("customes")
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

func (*customerRepository) Delete(Id uuid.UUID) error {
	_, err := DBInstance.Collection(_collectionName).DeleteOne(context.Background(), bson.M{"_id": Id})

	return err
}

func (*customerRepository) GetAll() (*[]response.CustomerResponseModel, error) {
	var customers *[]response.CustomerResponseModel

	cursor, err := DBInstance.Collection(_collectionName).Find(context.Background(), bson.M{})
	cursor.All(context.Background(), customers)

	return customers, err
}

func (*customerRepository) GetById(Id uuid.UUID) error {
	var customer *response.CustomerResponseModel
	return DBInstance.Collection(_collectionName).FindOne(context.Background(), bson.M{"_id": Id}).Decode(customer)
}
