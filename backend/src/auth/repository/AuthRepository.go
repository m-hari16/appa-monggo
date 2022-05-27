package repository

import (
	"go-fiber-app/src/auth/entity/domain"
	"go-fiber-app/src/auth/entity/request"
)

type AuthRepository interface {
	Find(req request.UserId) (err error, result domain.User)
	Login(request request.LoginRequest) (err error, result domain.User)
	Register(request domain.User) (err error, result domain.User)
	UpdateToken(email domain.Email, token domain.Token) (err error, result interface{})
}
