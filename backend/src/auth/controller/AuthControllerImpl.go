package controller

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &AuthControllerImpl{service: service}
}

func (a AuthControllerImpl) Register(c *fiber.Ctx) error {
	var req request.UserRequest
	err := c.BodyParser(&req)
	helper.PanicIfNeeded(err)

	code, service := a.service.Register(req)
	return c.Status(code).JSON(service)
}

func (a AuthControllerImpl) Login(c *fiber.Ctx) error {
	var req request.LoginRequest
	err := c.BodyParser(&req)
	helper.PanicIfNeeded(err)

	code, response := a.service.Login(req)
	return c.Status(code).JSON(response)
}

func (a AuthControllerImpl) Verify(c *fiber.Ctx) error {
	return c.JSON(helper.UnprocessableEntity())
}
