package user

type UserService interface {
	Register(user *User) error
}