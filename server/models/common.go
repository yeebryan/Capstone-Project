package models

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type State string

type Interval string

const (
	//cart states
	StateCompleted State = "completed"
	StateInProcess State = "inProcess"
	//playlist + order states
	StateOngoing   State = "ongoing"
	StatePaused    State = "paused"
	StatePending   State = "pending"
	StateDelivered State = "delivered"
	//shared states
	StateDeleted State = "deleted"
)

const (
	Weekly    Interval = "weekly"
	Bi_Weekly Interval = "biweekly"
	Monthly   Interval = "monthly"
)

func IntervalType(input string) (Interval, error) {
	switch strings.ToLower(input) {
	case "weekly":
		return Weekly, nil
	case "bi-weekly":
		return Bi_Weekly, nil
	case "monthly":
		return Monthly, nil
	}
	return "", errors.New("Interval not found")
}
