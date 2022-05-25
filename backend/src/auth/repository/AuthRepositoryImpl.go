package repository

import (
	"errors"
	"go-fiber-app/app"
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/domain"
	"go-fiber-app/src/auth/entity/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	collectionName = "user"
)

type AuthRepositoryImpl struct {
	db *mongo.Client
}

func NewAuthRepository(db *mongo.Client) AuthRepository {
	return AuthRepositoryImpl{db: db}
}

func (a AuthRepositoryImpl) Login(request request.LoginRequest) (err error, result domain.User) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := a.db.Database(app.GetDatabaseName()).Collection(collectionName)
	err = collection.FindOne(ctx, bson.M{"email": request.Email}).Decode(&result)

	if err != nil {
		return errors.New("Email or Password is wrong"), result
	}

	compare := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))
	if compare != nil {
		return errors.New("Email or Password is wrong"), result
	}

	return nil, result
}

func (a AuthRepositoryImpl) Register(request domain.User) (err error, result domain.User) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := a.db.Database(app.GetDatabaseName()).Collection(collectionName)

	_, err = collection.InsertOne(ctx, request)
	helper.PanicIfNeeded(err)
	if err != nil {
		errs := err.(mongo.WriteException)
		for _, e := range errs.WriteErrors {
			if e.Code == 11000 {
				return errors.New("Duplicate entry"), domain.User{}
			}
		}
	}

	return nil, request
}
