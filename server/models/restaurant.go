package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurant struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Address  string             `bson:"address,omitempty" json:"address,omitempty"`
	Category string             `bson:"category,omitempty" json:"category"`
	Image    *ImageData         `bson:"image,omitempty" json:"image,omitempty"`
	Menu     Menu               `bson:"menu,omitempty" json:"menu"` // list of food IDs
}

type Menu struct {
	Menu []primitive.ObjectID
}
