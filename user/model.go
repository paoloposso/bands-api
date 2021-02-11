package user

import (
	"errors"

	"github.com/gookit/validate"
)

type User struct {
	Name 		string `validate:"required|minLen:3"`
	Email 		string `validate:"email"`
	Password 	string `validate:"required|minLen:8"`
}

func (user User) ValidateRegister() error {
	v := validate.Struct(user)
	if !v.Validate() {
		return errors.New(v.Errors.String())
	}
	return nil
}