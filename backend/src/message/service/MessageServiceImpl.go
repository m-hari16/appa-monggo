package service

import (
	"go-fiber-app/app"
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	deviceRequest "go-fiber-app/src/device/entity/request"
	deviceRepository "go-fiber-app/src/device/repository"
	"go-fiber-app/src/message/entity/domain"
	"go-fiber-app/src/message/entity/request"
	"go-fiber-app/src/message/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageServiceImpl struct {
	repository repository.MessageRepository
	validate   *validator.Validate
	db         *mongo.Client
}

func (m MessageServiceImpl) Get(req authRequest.UserId) (httpCode int, response helper.Response) {
	err, result := m.repository.Get(req)

	if err != nil {
		return fiber.StatusInternalServerError, helper.ServerError(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}

func (m MessageServiceImpl) Create(req request.Message) (httpCode int, response helper.Response) {
	err := m.validate.Struct(req)
	if err != nil {
		return fiber.StatusBadRequest, helper.ErrValidate(err)
	}

	occuredAt := time.Now().Unix()
	messageLog := domain.MessageLog{
		Status:    domain.PendingStatus,
		OccuredAt: occuredAt,
	}

	deviceRepository := deviceRepository.NewDeviceRepository(app.NewDatabase().NewDB().(*mongo.Client))
	err, deviceResponse := deviceRepository.Find(deviceRequest.DeviceId(req.DeviceId))

	if err != nil {
		return fiber.StatusNotFound, helper.NotFound("Device not registered!")
	}

	deviceDomain := domain.Device{}
	copier.Copy(&deviceDomain, &deviceResponse)

	messageUser := domain.User{}
	copier.Copy(&messageUser, &deviceResponse.User)

	repoRequest := domain.Message{
		Id:          primitive.NewObjectID(),
		User:        messageUser,
		Device:      deviceDomain,
		PhoneNumber: req.PhoneNumber,
		Messages:    req.Messages,
		Status:      domain.PendingStatus,
		Log:         messageLog,
		CreatedAt:   occuredAt,
	}

	err, result := m.repository.Create(repoRequest)
	if err != nil {
		return fiber.StatusBadRequest, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}

func (m MessageServiceImpl) Show(req request.MessageId) (httpCode int, response helper.Response) {

	err, result := m.repository.Show(req)

	if err != nil {
		return fiber.StatusBadGateway, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}
