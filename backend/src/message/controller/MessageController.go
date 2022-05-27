package controller

import (
	"go-fiber-app/src/message/service"

	"github.com/gofiber/fiber/v2"
)

type MessageController interface {
	Create(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
}

func NewMessageController(service service.MessageService) MessageController {
	return MessageControllerImpl{service: service}
}
