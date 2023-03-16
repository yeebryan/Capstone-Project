package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Items      *[]FoodItems       `bson:"Items,omitempty" json:"Items"`
	TotalPrice float64            `bson:"total_price,omitempty" json:"total_price"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
