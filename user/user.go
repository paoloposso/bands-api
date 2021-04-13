package user

import (
	customerrors "github.com/paoloposso/bands-auth-api/custom_errors"

	"github.com/gookit/validate"
)

// User struct contains the User data
type User struct {
	ID			string `bson:"_id"`
	Name 		string `validate:"required|minLen:3"`
	Email 		string `validate:"required|email"`
	Password 	string `validate:"required|minLen:8"`
}

// ValidateRegister validates the User properties that are required for registering the User
func (user User) ValidateRegister() error {
	v := validate.Struct(user)
	if !v.Validate() {
		return &customerrors.DomainError { Message: v.Errors.String(), ErrorType: customerrors.InvalidDataError }
	}
	return nil
}