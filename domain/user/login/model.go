package login

import (
	customerrors "bands-api/custom_errors"

	"github.com/gookit/validate"
)

// Login contains user login data
type Login struct {
	Email		string	`validate:"required|email"`
	Password	string	`validate:"required"`
}

// ValidateLogin validates the user login data
func (l Login) ValidateLogin() error {
	v := validate.Struct(l)
	if !v.Validate() {
		return &customerrors.InvalidDataError { Message: v.Errors.String() }
	}
	return nil
}