package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PendingStatus MessageStatus = "pending"
	SendedStatus  MessageStatus = "sended"
)

type Message struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Device      Device             `json:"device" bson:"device"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Messages    string             `json:"messages" bson:"messages"`
	Status      MessageStatus      `json:"status" bson:"status"`
	Log         MessageLog         `json:"log" bson:"log"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
}

type MessageLog struct {
	Status    MessageStatus `json:"status" bson:"status"`
	OccuredAt int64         `json:"occured_at" bson:"occured_at"`
}

type Device struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Brand string             `json:"brand" bson:"brand"`
	Model string             `json:"model" bson:"model"`
}

type MessageStatus string

func (m MessageStatus) IsInvalid() error {
	switch m {
	case PendingStatus, SendedStatus:
		return nil
	}

	return errors.New("Invalid message status")
}
