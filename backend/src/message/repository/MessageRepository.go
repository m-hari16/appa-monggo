package repository

import (
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/message/entity/domain"
	"go-fiber-app/src/message/entity/request"

	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository interface {
	Get(req authRequest.UserId) (err error, result []domain.Message)
	Create(req domain.Message) (err error, result domain.Message)
	Show(req request.MessageId) (err error, result domain.Message)
	Update(messageId request.MessageId, messageLog domain.MessageLog, messageDevice domain.Device) (err error, result interface{})
}

func NewMessageRepository(db *mongo.Client) MessageRepository {
	return MessageRepositoryImpl{db: db}
}
