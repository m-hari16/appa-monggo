package domain

import (
	"errors"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @email field is unique
type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     Email              `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Token     Token              `json:"token" bson:"token"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at,omitempty"`
}

type Token string
type Email string

func (t Token) IsValid() error {
	if len(t) < 30 {
		return errors.New("Character token is to low")
	}

	return nil
}

func (e Email) IsValid() error {
	_, err := mail.ParseAddress(string(e))
	if err != nil {
		return errors.New("Email is not valid")
	}

	return nil
}
