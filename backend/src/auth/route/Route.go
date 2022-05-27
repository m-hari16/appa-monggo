package route

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/controller"
	"go-fiber-app/src/auth/repository"
	"go-fiber-app/src/auth/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func greeting(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helper.OnlyMessage("Auth module"))
}

func Register(route *fiber.App, db *mongo.Client) {
	repository := repository.NewAuthRepository(db)
	service := service.NewAuthService(repository, validator.New())
	controller := controller.NewAuthController(service)

	route.Post("/api/auth/register", controller.Register)
	route.Post("/api/auth/login", controller.Login)
	route.Patch("/api/auth/token/:email", controller.UpdateToken)
	route.Get("/api/auth/*", greeting)
}
