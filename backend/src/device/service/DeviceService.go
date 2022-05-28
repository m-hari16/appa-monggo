package service

import (
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/repository"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceService interface {
	Create(req request.Device) (httpCode int, response helper.Response)
	Get(req authRequest.UserId) (httpCode int, response helper.Response)
	Show(req request.DeviceId) (httpCode int, response helper.Response)
}

func NewDeviceService(repository repository.DeviceRepository, validate *validator.Validate, db *mongo.Client) DeviceService {
	return DeviceServiceImpl{repository: repository, validate: validate, db: db}
}
