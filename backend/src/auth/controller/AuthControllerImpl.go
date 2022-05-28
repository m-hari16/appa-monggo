package controller

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/domain"
	"go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/pkg"
	"go-fiber-app/src/auth/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	service service.AuthService
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

func (a AuthControllerImpl) UpdateToken(c *fiber.Ctx) error {
	data := pkg.NewJwtPkg()

	var userData domain.User
	helper.MapToStruct(data.GetTokenData(c), &userData)

	var email domain.Email
	email = domain.Email(c.Params("email"))

	if email != userData.Email {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.Unauthorized("Unauthorized"))
	}

	httpCode, response := a.service.UpdateToken(email)

	return c.Status(httpCode).JSON(response)
}
