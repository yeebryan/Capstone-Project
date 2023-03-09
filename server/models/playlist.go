package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	FoodID primitive.ObjectID `bson:"food_id,omitempty"`
	// UserID primitive.ObjectID `bson:"user_id,omitempty"`
	// Halal          bool               `bson:"halal,omitempty"`
	// Current        bool               `bson:"current,omitempty"`
	// Status         string             `bson:"status,omitempty"`
	// IntervalDates  time.Time          `bson:"interval_dates,omitempty"`
	// TimingInterval string             `bson:"timing_interval,omitempty"`
}
