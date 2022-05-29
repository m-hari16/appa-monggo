package service

import (
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/message/entity/domain"
	"go-fiber-app/src/message/entity/request"
	"go-fiber-app/src/message/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type MessageService interface {
	Get(req authRequest.UserId) (httpCode int, response helper.Response)
	Create(req request.Message, user domain.User) (httpCode int, response helper.Response)
	Show(req request.MessageId) (httpCode int, response helper.Response)
	Update(req request.MessageLogUpdate) (httpCode int, response helper.Response)
}

func NewMessageService(repository repository.MessageRepository, db *mongo.Client) MessageService {
	return MessageServiceImpl{repository: repository, db: db}
}
