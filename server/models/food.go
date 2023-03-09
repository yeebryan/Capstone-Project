package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	// Tag          string             `bson:"tag" json:"tag"`
	// Halal        bool               `bson:"halal" json:"halal"`
	RestaurantID primitive.ObjectID `bson:"restaurant_id" json:"restaurant_id"`
	// Image        *ImageData         `bson:"image,omitempty" json:"image,omitempty"`
}
