package pkg

import (
	"go-fiber-app/helper"
	"go-fiber-app/src/auth/entity/domain"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPkgImpl struct{}

func (j JwtPkgImpl) TokenClaims(user domain.User) (err error, t interface{}) {

	// Create the Claims
	claims := jwt.MapClaims{
		"data": user,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err, ""
	}
	return nil, t
}

func (j JwtPkgImpl) JwtWare() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: j.errorHandler,
	})
}

func (j JwtPkgImpl) errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.BadRequest(err.Error()))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(helper.Unauthorized("Unauthorize"))
}

func (j JwtPkgImpl) GetTokenData(c *fiber.Ctx) map[string]interface{} {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	data := claims["data"].(map[string]interface{})

	return data
}
