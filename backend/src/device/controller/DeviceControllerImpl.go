package controller

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/service"

	"github.com/gofiber/fiber/v2"
)

type DeviceControllerImpl struct {
	service service.DeviceService
}

func (d DeviceControllerImpl) Create(c *fiber.Ctx) error {
	var req request.Device
	err := c.BodyParser(&req)
	helper.PanicIfNeeded(err)

	httpCode, service := d.service.Create(req)

	return c.Status(httpCode).JSON(service)
}

func (d DeviceControllerImpl) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.StatusOK)
}

func (d DeviceControllerImpl) Show(c *fiber.Ctx) error {
	return c.JSON(fiber.StatusOK)
}

func (d DeviceControllerImpl) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.StatusOK)
}
