package service

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/device/entity/domain"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceServiceImpl struct {
	repository repository.DeviceRepository
	validate   *validator.Validate
}

func (d DeviceServiceImpl) Create(req request.Device) (httpCode int, response helper.Response) {
	err := d.validate.Struct(req)
	if err != nil {
		return fiber.StatusBadRequest, helper.ErrValidate(err)
	}

	svcReq := domain.Device{
		Id:          primitive.NewObjectID(),
		Brand:       req.Brand,
		Model:       req.Model,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   0,
	}

	err, repo := d.repository.Create(svcReq)
	if err != nil {
		return fiber.StatusBadRequest, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasStore(repo)
}

func (d DeviceServiceImpl) Get(req request.DeviceId) (httpCode int, response helper.Response) {
	return 0, helper.Response{}
}

func (d DeviceServiceImpl) Show(req request.DeviceId) (httpCode int, response helper.Response) {
	return 0, helper.Response{}
}

func (d DeviceServiceImpl) Delete(req request.DeviceId) (httpCode int, response helper.Response) {
	return 0, helper.Response{}
}
