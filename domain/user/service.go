package user

import "bands-api/domain/user/login"

// Service contains the definition of methods that execute the Business Logic for User Domain
type Service interface {
	Register(user *User) error
	Login(loginData login.Login) (string, error)
	GetDataByToken(token string) (*User, error)
}