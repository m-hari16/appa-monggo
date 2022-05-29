package controller

import (
	"go-fiber-app/src/auth/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Verify(c *fiber.Ctx) error
	UpdateToken(c *fiber.Ctx) error
}

func NewAuthController(service service.AuthService, validate *validator.Validate) AuthController {
	return &AuthControllerImpl{service: service, validate: validate}
}
