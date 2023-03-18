package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImageData struct {
	URL string `bson:"url" json:"url"`
	// FileName string `bson:"file_name" json:"file_name"`
}

type FoodItems struct {
	ID       primitive.ObjectID `bson:"food_id,omitempty" json:"food_id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Quantity int                `bson:"quantity,omitempty" json:"quantity"`
	Price    float64            `bson:"price,omitempty" json:"price"`
}
