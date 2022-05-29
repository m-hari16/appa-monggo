package main

import (
	"go-fiber-app/app"
	"go-fiber-app/helper"
	auth "go-fiber-app/src/auth/route"
	device "go-fiber-app/src/device/route"
	message "go-fiber-app/src/message/route"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func Greeting(c *fiber.Ctx) error {
	return c.JSON(helper.OnlyMessage("Restrict area"))
}

func main() {
	godotenv.Load(".env")
	database := app.NewDatabase().NewDB().(*mongo.Client)

	r := fiber.New()
	r.Get("/", Greeting)

	auth.Register(r, database)
	device.Register(r, database)
	message.Register(r, database)

	r.Listen(":" + os.Getenv("DB_PORT"))
}
