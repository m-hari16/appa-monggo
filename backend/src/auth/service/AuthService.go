package service

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/request"
)

type AuthService interface {
	Register(req request.UserRequest) (httpCode int, response helper.Response)
	Login(req request.LoginRequest) (httCode int, response helper.Response)
	Verify() (httpCode int, response bool)
}
