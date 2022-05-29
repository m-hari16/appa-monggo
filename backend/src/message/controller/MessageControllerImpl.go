package controller

import (
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/pkg"
	"go-fiber-app/src/message/entity/request"
	"go-fiber-app/src/message/service"

	"github.com/gofiber/fiber/v2"
)

type MessageControllerImpl struct {
	service service.MessageService
}

func (m MessageControllerImpl) Get(c *fiber.Ctx) error {
	jwt := pkg.NewJwtPkg()
	userData := jwt.GetTokenData(c)

	httpCode, response := m.service.Get(authRequest.UserId(userData["id"].(string)))

	return c.Status(httpCode).JSON(response)
}

func (m MessageControllerImpl) Create(c *fiber.Ctx) error {
	var req request.Message
	err := c.BodyParser(&req)
	helper.PanicIfNeeded(err)

	httpCode, response := m.service.Create(req)

	return c.Status(httpCode).JSON(response)
}

func (m MessageControllerImpl) Show(c *fiber.Ctx) error {
	var req request.MessageId
	req = request.MessageId(c.Params("message_id"))

	httpCode, response := m.service.Show(req)
	return c.Status(httpCode).JSON(response)
}

func (m MessageControllerImpl) Update(c *fiber.Ctx) error {

	var req request.MessageLogUpdate
	err := c.BodyParser(&req)
	helper.PanicIfNeeded(err)

	httpCode, response := m.service.Update(req)

	return c.Status(httpCode).JSON(response)
}
