package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Items         *[]FoodItems       `bson:"Items,omitempty" json:"Items"`
	TotalPrice    float64            `bson:"total_price,omitempty" json:"total_price"`
	PaymentMethod string             `bson:"payment_method,omitempty" json:"payment_method"`
	DeliveryTime  time.Time          `bson:"delivery_time,omitempty" json:"delivery_time"`
	CreatedAt     time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
