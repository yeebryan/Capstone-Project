package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	Price       float64            `bson:"price,omitempty" json:"price,omitempty"`
	Image       *ImageData         `bson:"image,omitempty" json:"image,omitempty"`
	Tag         []string           `bson:"tag,omitempty" json:"tag,omitempty"`
}
