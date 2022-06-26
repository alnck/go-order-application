package repository

import (
	"context"
	"customer-service/src/domain/entity"
	"customer-service/src/infrastructure/interfaces"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _collectionName = "customers"

var lock sync.Mutex

type customerRepository struct{}

var DBInstance *mongo.Database

func initDBInstance() {
	if DBInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if DBInstance == nil {
			clientOptions := options.Client().ApplyURI("mongodb://c_mongodb:27017") //LINK_MONGODB_DOCKER_URI
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
			DBInstance = client.Database("customerDB")
		}
	}
}

func NewCustomerRepository() interfaces.ICustomerRepository {
	initDBInstance()

	return &customerRepository{}
}

func (*customerRepository) Create(customer *entity.Customer) (interface{}, error) {
	result, err := DBInstance.Collection(_collectionName).
		InsertOne(
			context.Background(),
			&customer,
			options.InsertOne())

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (*customerRepository) Update(customer *entity.Customer) (bool, error) {
	updateDocument := bson.M{
		"$set": bson.M{
			"Name":      customer.Name,
			"Email":     customer.Email,
			"Address":   customer.Address,
			"UpdatedAt": time.Now().UTC(),
		},
	}

	result, err := DBInstance.Collection(_collectionName).
		UpdateOne(
			context.Background(),
			bson.M{"_id": customer.Id},
			updateDocument)

	if err != nil {
		return false, err
	}

	return result.ModifiedCount > 0, nil
}

func (*customerRepository) Delete(Id primitive.ObjectID) (bool, error) {
	result, err := DBInstance.Collection(_collectionName).DeleteOne(context.Background(), bson.M{"_id": Id})

	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}

func (*customerRepository) GetAll(page int, limit int) ([]entity.Customer, error) {
	options := new(options.FindOptions)

	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := DBInstance.Collection(_collectionName).Find(context.Background(), bson.M{}, options)

	if err != nil {
		return nil, err
	}

	var customer []entity.Customer

	err = cursor.All(context.Background(), &customer)

	return customer, err
}

func (*customerRepository) GetById(Id primitive.ObjectID) (entity.Customer, error) {
	var customer entity.Customer

	err := DBInstance.Collection(_collectionName).FindOne(context.Background(), bson.M{"_id": Id}).Decode(&customer)

	return customer, err
}

func (*customerRepository) IsValid(Id primitive.ObjectID) (bool, error) {
	count, err := DBInstance.Collection(_collectionName).CountDocuments(context.Background(), bson.M{"_id": Id})
	check := count > 0
	return check, err
}
