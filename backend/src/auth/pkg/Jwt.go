package pkg

import (
	"go-fiber-app/src/auth/entity/domain"

	"github.com/gofiber/fiber/v2"
)

type JwtPkg interface {
	TokenClaims(user domain.User) (err error, token interface{})
	JwtWare() fiber.Handler
	errorHandler(c *fiber.Ctx, err error) error
	GetTokenData(c *fiber.Ctx) map[string]interface{}
}

func NewJwtPkg() JwtPkg {
	return JwtPkgImpl{}
}
