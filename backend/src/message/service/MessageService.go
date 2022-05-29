package service

import (
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/message/entity/request"
	"go-fiber-app/src/message/repository"

	"github.com/go-playground/validator/v10"
)

type MessageService interface {
	Get(req authRequest.UserId) (httpCode int, response helper.Response)
	Create(req request.Message) (httpCode int, response helper.Response)
	Show(req request.MessageId) (httpCode int, response helper.Response)
	Update(req request.MessageLogUpdate) (httpCode int, response helper.Response)
}

func NewMessageService(repository repository.MessageRepository, validate *validator.Validate) MessageService {
	return MessageServiceImpl{repository: repository, validate: validate}
}
