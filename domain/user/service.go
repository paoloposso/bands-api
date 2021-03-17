package user

// Service contains the definition of methods that execute the Business Logic for User Domain
type Service interface {
	Register(user *User) error
	Login(email string, password string) (string, error)
	GetDataByToken(token string) (*User, error)
}