package repository

import (
	"go-fiber-app/src/message/entity/domain"
	"go-fiber-app/src/message/entity/request"

	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository interface {
	Create(req domain.Message) (err error, result domain.Message)
	Show(req request.MessageId) (err error, result domain.Message)
}

func NewMessageRepository(db *mongo.Client) MessageRepository {
	return MessageRepositoryImpl{db: db}
}
