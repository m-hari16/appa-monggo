package service

import (
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceService interface {
	Create(req request.Device) (httpCode int, response helper.Response)
	Get(req authRequest.UserId) (httpCode int, response helper.Response)
	Show(req request.DeviceId) (httpCode int, response helper.Response)
}

func NewDeviceService(repository repository.DeviceRepository, db *mongo.Client) DeviceService {
	return DeviceServiceImpl{repository: repository, db: db}
}
