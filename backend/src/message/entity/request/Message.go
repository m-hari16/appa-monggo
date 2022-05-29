package request

import (
	deviceRequest "go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/message/entity/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id          primitive.ObjectID     `json:"id"`
	DeviceId    deviceRequest.DeviceId `json:"device_id"`
	PhoneNumber string                 `json:"phone_number" validate:"required"`
	Messages    string                 `json:"messages" validate:"required"`
}

type MessageLogUpdate struct {
	Id        MessageId              `json:"id" validate:"required"`
	DeviceId  deviceRequest.DeviceId `json:"device_id" validate:"required"`
	Status    domain.MessageStatus   `json:"status" validate:"required"`
	OccuredAt int64                  `json:"occured_at"`
}

type MessageId string
