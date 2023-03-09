package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurant struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CuisineType    string             `bson:"cuisine_type,omitempty" json:"cuisine_type"`
	RestaurantName string             `bson:"restaurant_name,omitempty" json:"restaurant_name"`
	// Rating         float32            `bson:"rating" json:"rating"`
	// Location       string             `bson:"location" json:"location"`
	// PostalCode     int                `bson:"postal_code" json:"postal_code"`
	// OpeningHours   []string           `bson:"opening_hours,omitempty" json:"opening_hours,omitempty"`
	// WeeklySchedule string             `bson:"weekly_schedule,omitempty" json:"weekly_schedule,omitempty"`
	// Reviews        string             `bson:"reviews,omitempty" json:"reviews,omitempty"`
	// Image          *ImageData         `bson:"image,omitempty" json:"image,omitempty"`
}
