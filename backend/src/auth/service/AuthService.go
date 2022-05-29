package service

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/domain"
	"go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/repository"
)

type AuthService interface {
	Register(req request.UserRequest) (httpCode int, response helper.Response)
	Login(req request.LoginRequest) (httCode int, response helper.Response)
	Verify() (httpCode int, response bool)
	UpdateToken(email domain.Email) (httpCode int, response helper.Response)
}

func NewAuthService(repository repository.AuthRepository) AuthService {
	return AuthServiceImpl{repository: repository}
}
