package repository

import (
	"errors"
	"fmt"
	"go-fiber-app/app"
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
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

func (m MessageRepositoryImpl) Get(req authRequest.UserId) (err error, result []domain.Message) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	obj, _ := primitive.ObjectIDFromHex(string(req))
	filter := bson.M{"user._id": obj}

	collection := m.db.Database(app.GetDatabaseName()).Collection(collectionName)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return err, result
	}

	for cursor.Next(ctx) {
		var tmpResult domain.Message
		err = cursor.Decode(&tmpResult)
		helper.PanicIfNeeded(err)

		result = append(result, tmpResult)
	}

	return nil, result
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

func (m MessageRepositoryImpl) Update(messageId request.MessageId, req domain.MessageLog) (err error, result interface{}) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	err = req.Status.IsInvalid()
	if err != nil {
		return err, result
	}

	id, _ := primitive.ObjectIDFromHex(string(messageId))
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"log":    req,
		"status": req.Status,
	}}

	collection := m.db.Database(app.GetDatabaseName()).Collection(collectionName)
	result, err = collection.UpdateOne(ctx, filter, update)
	tmp := result.(*mongo.UpdateResult)

	if err != nil {
		return err, result
	}

	if tmp.MatchedCount < 1 {
		fmt.Println("no")
		return errors.New("No data found"), nil
	}

	return nil, result
}
