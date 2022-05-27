package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id          primitive.ObjectID `json:"id"`
	DeviceId    string             `json:"device_id" validate:"required"`
	PhoneNumber string             `json:"phone_number" validate:"required"`
	Messages    string             `json:"messages" validate:"required"`
}

type MessageId string
