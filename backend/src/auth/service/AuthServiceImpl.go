package service

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/domain"
	"go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repository repository.AuthRepository
	validate   *validator.Validate
}

func NewAuthService(repository repository.AuthRepository, validate *validator.Validate) AuthService {
	return AuthServiceImpl{repository: repository, validate: validate}
}

func (a AuthServiceImpl) Login(req request.LoginRequest) (httpCode int, response interface{}) {
	err := a.validate.Struct(req)
	if err != nil {
		return fiber.StatusBadRequest, helper.ErrValidate(err)
	}

	err, repo := a.repository.Login(req)

	if err != nil {
		return fiber.StatusNotFound, helper.Unauthorized(err.Error())
	}

	return fiber.StatusOK, helper.HasOk(repo)
}

func (a AuthServiceImpl) Register(req request.UserRequest) (httpCode int, response interface{}) {
	err := a.validate.Struct(req)
	if err != nil {
		return fiber.StatusBadRequest, helper.ErrValidate(err)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	svcReq := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPass),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: 0,
	}

	err, repo := a.repository.Register(svcReq)
	if err != nil {
		return fiber.StatusUnprocessableEntity, helper.BadRequest(err.Error())
	}

	return fiber.StatusOK, helper.HasStore(repo)
}

func (a AuthServiceImpl) Verify() (httpCode int, response bool) {
	return fiber.StatusOK, true
}
