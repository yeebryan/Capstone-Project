package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name      string               `bson:"name,omitempty"`
	FoodID    []primitive.ObjectID `bson:"food_id,omitempty" json:"food_id"`
	UserID    primitive.ObjectID   `bson:"user_id,omitempty" json:"user_id"`
	Image     *ImageData           `bson:"image,omitempty" json:"image,omitempty"`
	Status    string               `bson:"status,omitempty" json:"status,omitempty"`
	StartDate string               `bson:"start_date,omitempty" json:"start_date,omitempty"`
	// Halal          bool               `bson:"halal,omitempty"`
	// Current        bool               `bson:"current,omitempty"`
	// IntervalDates  time.Time          `bson:"interval_dates,omitempty"`
	// TimingInterval string             `bson:"timing_interval,omitempty"`
}
