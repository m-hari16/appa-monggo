package repository

import (
	"errors"
	"fmt"
	"go-fiber-app/app"
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/domain"
	"go-fiber-app/src/auth/entity/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (a AuthRepositoryImpl) Find(req request.UserId) (err error, result domain.User) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := a.db.Database(app.GetDatabaseName()).Collection(collectionName)
	userId, _ := primitive.ObjectIDFromHex(string(req))
	err = collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return err, result
	}

	return nil, result
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

	if request.Token.IsValid() != nil {
		return request.Token.IsValid(), domain.User{}
	}

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

func (a AuthRepositoryImpl) UpdateToken(email domain.Email, token domain.Token) (err error, result interface{}) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	filter := bson.M{
		"email": bson.M{
			"$eq": email,
		},
	}

	update := bson.M{
		"$set": bson.M{"token": token},
	}

	collection := a.db.Database(app.GetDatabaseName()).Collection(collectionName)
	result, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Update error"), domain.User{}
	}

	return nil, result
}
