package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @model field is unique
type Device struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MacAddress  string             `json:"mac_address" bson:"mac_address"`
	User        User               `json:"user" bson:"user"`
	Brand       string             `json:"brand" bson:"brand"`
	Model       string             `json:"model" bson:"model"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
	UpdatedAt   int64              `json:"updated_at" bson:"updated_at"`
}

type User struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}
