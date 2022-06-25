package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id       primitive.ObjectID `json:"Id" bson:"_id"`
	ImageUrl string             `json:"ImageUrl" bson:"ImageUrl"`
	Name     string             `json:"Name" bson:"Name"`
}
