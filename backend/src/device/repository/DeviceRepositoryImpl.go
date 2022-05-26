package repository

import (
	"errors"
	"fmt"
	"go-fiber-app/app"
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/device/entity/domain"
	"go-fiber-app/src/device/entity/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "device"
)

type DeviceRepositoryImpl struct {
	Db *mongo.Client
}

func (d DeviceRepositoryImpl) Create(req domain.Device) (err error, result domain.Device) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := d.Db.Database(app.GetDatabaseName()).Collection(collectionName)

	_, err = collection.InsertOne(ctx, req)
	if err != nil {
		errs := err.(mongo.WriteException)
		for _, e := range errs.WriteErrors {
			if e.Code == 11000 {
				return errors.New("Duplicate entry"), domain.Device{}
			}
		}
	}

	return nil, req
}

func (d DeviceRepositoryImpl) Get(req authRequest.UserId) (err error, result []domain.Device) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := d.Db.Database(app.GetDatabaseName()).Collection(collectionName)
	obj, _ := primitive.ObjectIDFromHex(string(req))
	cursor, err := collection.Find(ctx, bson.M{"_id": obj})
	if err != nil {
		fmt.Println(errors.New(err.Error()))
		return err, result
	}

	for cursor.Next(ctx) {
		var tmpResult domain.Device
		err := cursor.Decode(&tmpResult)
		helper.PanicIfNeeded(err)

		result = append(result, tmpResult)
	}

	return nil, result
}

func (d DeviceRepositoryImpl) Find(req request.DeviceId) (err error, result domain.Device) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := d.Db.Database(app.GetDatabaseName()).Collection(collectionName)
	obj, _ := primitive.ObjectIDFromHex(string(req))
	err = collection.FindOne(ctx, bson.M{"_id": obj}).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Data not found"), result
	}

	return nil, result
}

func (d DeviceRepositoryImpl) Delete(req request.DeviceId) (err error, result request.DeviceId) {
	return nil, result
}
