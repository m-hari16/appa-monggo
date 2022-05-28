package controller

import (
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/pkg"
	"go-fiber-app/src/device/entity/request"
	"go-fiber-app/src/device/service"

	"github.com/gofiber/fiber/v2"
)

type DeviceControllerImpl struct {
	service service.DeviceService
}

func (d DeviceControllerImpl) Create(c *fiber.Ctx) error {
	var req request.Device
	_ = c.BodyParser(&req)

	jwt := pkg.NewJwtPkg()
	userData := jwt.GetTokenData(c)
	req.UserId = userData["id"].(string)

	httpCode, response := d.service.Create(req)

	return c.Status(httpCode).JSON(response)
}

func (d DeviceControllerImpl) Get(c *fiber.Ctx) error {
	var userId authRequest.UserId
	jwt := pkg.NewJwtPkg()
	userData := jwt.GetTokenData(c)
	userId = authRequest.UserId(userData["id"].(string))

	httpCode, response := d.service.Get(userId)

	return c.Status(httpCode).JSON(response)
}

func (d DeviceControllerImpl) Show(c *fiber.Ctx) error {
	deviceId := request.DeviceId(c.Params("device_id"))

	httpCode, response := d.service.Show(deviceId)

	return c.Status(httpCode).JSON(response)
}
