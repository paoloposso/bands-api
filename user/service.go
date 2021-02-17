package user

type UserService interface {
	Register(user *User) error
	Login(email string, password string) (string, error)
}