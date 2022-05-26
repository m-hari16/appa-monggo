package service

import (
	"go-fiber-app/helper"
	authReq "go-fiber-app/src/auth/entity/request"
	authRepo "go-fiber-app/src/auth/repository"
	"go-fiber-app/src/device/entity/domain"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceServiceImpl struct {
	repository repository.DeviceRepository
	validate   *validator.Validate
	db         *mongo.Client
}

func (d DeviceServiceImpl) Create(req request.Device) (httpCode int, response helper.Response) {
	err := d.validate.Struct(req)
	if err != nil {
		return fiber.StatusBadRequest, helper.ErrValidate(err)
	}

	authRepository := authRepo.NewAuthRepository(d.db)
	err, authResponse := authRepository.Find(authReq.UserId(req.UserId))
	userDomain := domain.User{}
	copier.Copy(&userDomain, &authResponse)

	repoRequest := domain.Device{
		Id:          primitive.NewObjectID(),
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

func (d DeviceServiceImpl) Get() (httpCode int, response helper.Response) {
	/** getting userid from jwt */

	return fiber.StatusOK, helper.HasOk(nil)
}

func (d DeviceServiceImpl) Show(req request.DeviceId) (httpCode int, response helper.Response) {
	err, result := d.repository.Find(req)

	if err != nil {
		return fiber.StatusNotFound, helper.NotFound(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}
