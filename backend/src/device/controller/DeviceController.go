package controller

import (
	"go-fiber-app/src/device/service"

	"github.com/gofiber/fiber/v2"
)

type DeviceController interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
}

func NewDeviceController(service service.DeviceService) DeviceController {
	return DeviceControllerImpl{service: service}
}
