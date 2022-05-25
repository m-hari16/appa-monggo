package helper

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

func PanicIfNeeded(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ErrValidate(message error) Response {
	messages := TranslateError(message)
	return Response{Code: strconv.Itoa(fiber.StatusUnprocessableEntity), Success: true, Message: messages}
}

func TranslateError(err error) (errs []string) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		errs = append(errs, ToSnake(e.Field())+" is "+e.Tag())
	}
	return errs
}
