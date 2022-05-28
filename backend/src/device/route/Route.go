package route

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/pkg"
	"go-fiber-app/src/device/controller"
	"go-fiber-app/src/device/repository"
	"go-fiber-app/src/device/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func greeting(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helper.OnlyMessage("Device module"))
}

func Register(route *fiber.App, db *mongo.Client) {
	repository := repository.NewDeviceRepository(db)
	service := service.NewDeviceService(repository, validator.New(), db)
	controller := controller.NewDeviceController(service)
	jwt := pkg.NewJwtPkg()

	route.Get("/api/device/greeting", greeting)
	jwt.JwtWare(route)
	route.Get("/api/device", controller.Get)
	route.Get("/api/device/:device_id", controller.Show)
	route.Post("/api/device", controller.Create)
}
