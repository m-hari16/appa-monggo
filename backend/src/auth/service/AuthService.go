package service

import (
	"go-fiber-app/src/auth/entity/request"
)

type AuthService interface {
	Register(req request.UserRequest) (httpCode int, response interface{})
	Login(req request.LoginRequest) (httCode int, response interface{})
	Verify() (httpCode int, response bool)
}
