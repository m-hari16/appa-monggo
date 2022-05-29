package service

import (
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	authRepository "go-fiber-app/src/auth/repository"
	"go-fiber-app/src/device/entity/domain"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceServiceImpl struct {
	repository repository.DeviceRepository
	db         *mongo.Client
}

func (d DeviceServiceImpl) Create(req request.Device) (httpCode int, response helper.Response) {

	authRepository := authRepository.NewAuthRepository(d.db)
	err, authResponse := authRepository.Find(authRequest.UserId(req.UserId))
	userDomain := domain.User{}
	copier.Copy(&userDomain, &authResponse)

	repoRequest := domain.Device{
		Id:          primitive.NewObjectID(),
		MacAddress:  req.MacAddress,
		User:        userDomain,
		Brand:       req.Brand,
		Model:       req.Model,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   0,
	}

	err, result := d.repository.Create(repoRequest)
	if err != nil {
		return fiber.StatusBadRequest, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasStore(result)
}

func (d DeviceServiceImpl) Get(req authRequest.UserId) (httpCode int, response helper.Response) {
	_, result := d.repository.Get(req)

	return fiber.StatusOK, helper.HasOk(result)
}

func (d DeviceServiceImpl) Show(req request.DeviceId) (httpCode int, response helper.Response) {
	err, result := d.repository.Find(req)

	if err != nil {
		return fiber.StatusNotFound, helper.NotFound(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}
