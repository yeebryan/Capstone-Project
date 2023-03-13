package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"username"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Password    string             `bson:"password,omitempty" json:"-"`
	Address     string             `bson:"address,omitempty" json:"address,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}
