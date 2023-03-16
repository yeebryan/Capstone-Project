package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Email        *string            `bson:"email" json:"email" validate:"email,required"`
	Password     *string            `bson:"password" json:"password" validate:"required,min=6"`
	Address      string             `bson:"address,omitempty" json:"address,omitempty"`
	PhoneNumber  *string            `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Token        *string            `bson:"token,omitempty" json:"token,omitempty"`
	RefreshToken *string            `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	UserID       string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
}
