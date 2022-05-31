package helper

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
)

type StructTest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func TestToSnake(t *testing.T) {
	text := "thisIsPascalCase"
	result := ToSnake(text)

	if result == "this_is_pascal_case" {
		return
	}

	t.Error("to snake case is not as expected")
}

func TestMapToStruct(t *testing.T) {
	data := make(map[string]interface{})
	data["first_name"] = "Zidni"
	data["last_name"] = "Mujib"

	result := StructTest{}

	MapToStruct(data, &result)

	if result.FirstName == "Zidni" && result.LastName == "Mujib" {
		return
	}

	t.Error("Expected must be struct must be filled")
}

func TestTranslateErr(t *testing.T) {
	data := StructTest{FirstName: "Zidni"}
	validate := validator.New()
	err := validate.Struct(data)

	result := TranslateError(err)
	var expected []string
	expected = []string{"last_name is required"}

	if result[0] == expected[0] && reflect.TypeOf(result) == reflect.TypeOf(expected) {
		return
	}

	t.Error("Expected is `last_name is required` type of []string")
}

func TestGetRandomString(t *testing.T) {
	data := RandomString(10)
	var expected string

	if len(data) == 10 && reflect.TypeOf(data) == reflect.TypeOf(expected) {
		return
	}

	t.Error("Expected is random string with 10 digit but", data)
}
