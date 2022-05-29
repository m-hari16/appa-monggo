package controller

import (
	"go-fiber-app/src/message/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageController interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

func NewMessageController(service service.MessageService, validate *validator.Validate, db *mongo.Client) MessageController {
	return MessageControllerImpl{service: service, validate: validate, db: db}
}
