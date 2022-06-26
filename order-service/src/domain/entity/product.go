package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id       primitive.ObjectID `json:"Id" bson:"_id" validate:"required"`
	ImageUrl string             `json:"ImageUrl" bson:"ImageUrl" validate:"required"`
	Name     string             `json:"Name" bson:"Name" validate:"required"`
}
