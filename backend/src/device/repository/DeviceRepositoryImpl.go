package repository

import (
	"errors"
	"go-fiber-app/app"
	"go-fiber-app/helper"
	"go-fiber-app/src/device/entity/domain"
	"go-fiber-app/src/device/entity/request"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "device"
)

type DeviceRepositoryImpl struct {
	db *mongo.Client
}

func (d DeviceRepositoryImpl) Create(req domain.Device) (err error, result domain.Device) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := d.db.Database(app.GetDatabaseName()).Collection(collectionName)

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

func (d DeviceRepositoryImpl) Find(req request.DeviceId) (err error, result domain.Device) {
	return nil, result
}

func (d DeviceRepositoryImpl) Delete(req request.DeviceId) (err error, result request.DeviceId) {
	return nil, result
}
