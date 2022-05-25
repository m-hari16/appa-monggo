package repository

import (
	"go-fiber-app/src/device/entity/domain"
	"go-fiber-app/src/device/entity/request"

	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceRepository interface {
	Create(req domain.Device) (err error, result domain.Device)
	Find(req request.DeviceId) (err error, result domain.Device)
	Delete(req request.DeviceId) (err error, result request.DeviceId)
}

func NewDeviceRepository(db *mongo.Client) DeviceRepository {
	return DeviceRepositoryImpl{db: db}
}
