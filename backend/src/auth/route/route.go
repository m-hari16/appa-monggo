package route

import (
	"go-fiber-app/src/auth/controller"
	"go-fiber-app/src/auth/repository"
	"go-fiber-app/src/auth/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Greeting(c *fiber.Ctx) error {
	return c.SendString("Auth")
}

func Register(route *fiber.App, db *mongo.Client) {
	repository := repository.NewAuthRepository(db)
	service := service.NewAuthService(repository, validator.New())
	controller := controller.NewAuthController(service)

	route.Post("/api/register", controller.Register)
	route.Post("/api/login", controller.Login)
}
