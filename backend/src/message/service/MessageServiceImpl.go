package service

import (
	"errors"
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	deviceRepository "go-fiber-app/src/device/repository"
	"go-fiber-app/src/message/entity/domain"
	"go-fiber-app/src/message/entity/request"
	"go-fiber-app/src/message/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageServiceImpl struct {
	repository repository.MessageRepository
	db         *mongo.Client
}

func (m MessageServiceImpl) Get(req authRequest.UserId) (httpCode int, response helper.Response) {
	err, result := m.repository.Get(req)

	if err != nil {
		return fiber.StatusInternalServerError, helper.ServerError(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}

func (m MessageServiceImpl) Create(req request.Message, user domain.User) (httpCode int, response helper.Response) {

	occuredAt := time.Now().Unix()
	messageLog := domain.MessageLog{
		Status:    domain.PendingStatus,
		OccuredAt: occuredAt,
	}

	repoRequest := domain.Message{
		Id:          primitive.NewObjectID(),
		User:        user,
		Device:      domain.Device{Id: primitive.NewObjectID()},
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

func (m MessageServiceImpl) Update(req request.MessageLogUpdate) (httpCode int, reponse helper.Response) {
	var device domain.Device
	messageLog := domain.MessageLog{
		Status:    req.Status,
		OccuredAt: time.Now().Unix(),
	}

	deviceRepository := deviceRepository.NewDeviceRepository(m.db)
	err, deviceResponse := deviceRepository.Find(req.DeviceId)

	if err != nil {
		return fiber.StatusNotFound, helper.NotFound(errors.New("Device not registered").Error())
	}

	copier.Copy(&device, &deviceResponse)

	err, result := m.repository.Update(req.Id, messageLog, device)

	if err != nil {
		return fiber.StatusBadRequest, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}
