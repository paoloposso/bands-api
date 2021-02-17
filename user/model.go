package user

import (
	customerrors "bands-api/custom_errors"

	"github.com/gookit/validate"
)

// User struct contains the User data
type User struct {
	ID			string
	Name 		string `validate:"required|minLen:3"`
	Email 		string `validate:"email"`
	Password 	string `validate:"required|minLen:8"`
}

// ValidateRegister validates the User properties that are required for registering the User
func (user User) ValidateRegister() error {
	v := validate.Struct(user)
	if !v.Validate() {
		return &customerrors.InvalidDataError { Message: v.Errors.String() }
	}
	return nil
}