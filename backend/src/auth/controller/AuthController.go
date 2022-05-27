package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Verify(c *fiber.Ctx) error
	UpdateToken(c *fiber.Ctx) error
}
