package controller

import (
	"go-fiber-app/helper"
	authRequest "go-fiber-app/src/auth/entity/request"
	"go-fiber-app/src/auth/pkg"
	authRepository "go-fiber-app/src/auth/repository"
	"go-fiber-app/src/message/entity/domain"
	"go-fiber-app/src/message/entity/request"
	"go-fiber-app/src/message/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageControllerImpl struct {
	service  service.MessageService
	validate *validator.Validate
	db       *mongo.Client
}

func (m MessageControllerImpl) Get(c *fiber.Ctx) error {
	jwt := pkg.NewJwtPkg()
	userData := jwt.GetTokenData(c)

	httpCode, response := m.service.Get(authRequest.UserId(userData["id"].(string)))

	return c.Status(httpCode).JSON(response)
}

func (m MessageControllerImpl) Create(c *fiber.Ctx) error {
	var req request.Message
	var user domain.User

	_ = c.BodyParser(&req)
	userToken := string(c.Request().Header.Peek("x-api-key"))

	userRepository := authRepository.NewAuthRepository(m.db)
	err, userResponse := userRepository.FindByToken(userToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.NotFound("Unauthorized"))
	}
	copier.Copy(&user, &userResponse)

	err = m.validate.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrValidate(err))
	}

	if string(userResponse.Token) != userToken {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.Unauthorized("Unauthorized"))
	}

	httpCode, response := m.service.Create(req, user)

	return c.Status(httpCode).JSON(response)
}

func (m MessageControllerImpl) Show(c *fiber.Ctx) error {
	var req request.MessageId
	req = request.MessageId(c.Params("message_id"))

	httpCode, response := m.service.Show(req)
	return c.Status(httpCode).JSON(response)
}

func (m MessageControllerImpl) Update(c *fiber.Ctx) error {
	var req request.MessageLogUpdate
	err := c.BodyParser(&req)

	err = m.validate.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrValidate(err))
	}

	httpCode, response := m.service.Update(req)

	return c.Status(httpCode).JSON(response)
}
