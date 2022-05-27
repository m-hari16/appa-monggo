package repository

import (
	"errors"
	"go-fiber-app/app"
	"go-fiber-app/helper"
	"go-fiber-app/src/message/entity/domain"
	"go-fiber-app/src/message/entity/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "messages"
)

type MessageRepositoryImpl struct {
	db *mongo.Client
}

func (m MessageRepositoryImpl) Create(req domain.Message) (err error, result domain.Message) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := m.db.Database(app.GetDatabaseName()).Collection(collectionName)
	_, err = collection.InsertOne(ctx, req)
	if err != nil {
		errs := err.(mongo.WriteException)
		for _, e := range errs.WriteErrors {
			if e.Code == 11000 {
				return errors.New("Duplicate entry"), domain.Message{}
			}
		}
	}

	return nil, req
}

func (m MessageRepositoryImpl) Show(req request.MessageId) (err error, result domain.Message) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	messageId, _ := primitive.ObjectIDFromHex(string(req))
	collection := m.db.Database(app.GetDatabaseName()).Collection(collectionName)
	err = collection.FindOne(ctx, bson.M{"_id": messageId}).Decode(&result)

	if err != nil {
		return err, domain.Message{}
	}

	return nil, result
}
