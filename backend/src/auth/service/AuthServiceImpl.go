package service

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/domain"
	"go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/pkg"
	"go-fiber-app/src/auth/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repository repository.AuthRepository
	validate   *validator.Validate
}

func NewAuthService(repository repository.AuthRepository, validate *validator.Validate) AuthService {
	return AuthServiceImpl{repository: repository, validate: validate}
}

func (a AuthServiceImpl) Login(req request.LoginRequest) (httpCode int, response helper.Response) {
	err := a.validate.Struct(req)
	if err != nil {
		return fiber.StatusBadRequest, helper.ErrValidate(err)
	}

	err, result := a.repository.Login(req)

	if err != nil {
		return fiber.StatusNotFound, helper.Unauthorized(err.Error())
	}

	jwt := pkg.NewJwtPkg()
	err, token := jwt.TokenClaims(result)

	return fiber.StatusOK, helper.HasOk(fiber.Map{"access_token": token, "user": result})
}

func (a AuthServiceImpl) Register(req request.UserRequest) (httpCode int, response helper.Response) {
	err := a.validate.Struct(req)
	if err != nil {
		return fiber.StatusBadRequest, helper.ErrValidate(err)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	repoRequest := domain.User{
		Id:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     domain.Email(req.Email),
		Password:  string(hashedPass),
		CreatedAt: time.Now().Unix(),
		Token:     domain.Token(helper.RandomString(30)),
		UpdatedAt: 0,
	}

	err, result := a.repository.Register(repoRequest)
	if err != nil {
		return fiber.StatusUnprocessableEntity, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasStore(result)
}

func (a AuthServiceImpl) UpdateToken(email domain.Email) (httpCode int, response helper.Response) {
	err, result := a.repository.UpdateToken(email, domain.Token(helper.RandomString(30)))

	if err != nil {
		return fiber.StatusUnprocessableEntity, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(result)
}

func (a AuthServiceImpl) Verify() (httpCode int, response bool) {
	return fiber.StatusOK, true
}
