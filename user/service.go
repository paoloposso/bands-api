package user

import "github.com/dgrijalva/jwt-go"

// Service contains the definition of methods that execute the Business Logic for User Domain
type Service interface {
	Register(user *User) error
	Login(email string, password string) (string, error)
	CheckLoginWithToken(tokenString string) (*jwt.Token, error)
}