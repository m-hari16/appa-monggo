package service

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/repository"

	"github.com/go-playground/validator/v10"
)

type DeviceService interface {
	Create(req request.Device) (httpCode int, response helper.Response)
	Get(req request.DeviceId) (httpCode int, response helper.Response)
	Show(req request.DeviceId) (httpCode int, response helper.Response)
	Delete(req request.DeviceId) (httpCode int, response helper.Response)
}

func NewDeviceService(repository repository.DeviceRepository, validate *validator.Validate) DeviceService {
	return DeviceServiceImpl{repository: repository, validate: validate}
}
