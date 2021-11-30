package model

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
