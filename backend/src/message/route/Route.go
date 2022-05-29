package route

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/pkg"
	"go-fiber-app/src/message/controller"
	"go-fiber-app/src/message/repository"
	"go-fiber-app/src/message/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func greeting(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helper.OnlyMessage("Message module"))
}

func Register(route *fiber.App, db *mongo.Client) {
	repository := repository.NewMessageRepository(db)
	service := service.NewMessageService(repository, validator.New())
	controller := controller.NewMessageController(service)
	jwt := pkg.NewJwtPkg()

	route.Get("/api/message/greeting", greeting)
	jwt.JwtWare(route)
	route.Get("/api/message", controller.Get)
	route.Get("/api/message/:message_id", controller.Show)
	route.Post("/api/message", controller.Create)
	route.Patch("api/message", controller.Update)
}
