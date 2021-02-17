package memory

import (
	"bands-api/user"
)

type userRepository struct {}

// NewMemoryRepository returns a reference to an implementation of UserRepository interface
func NewMemoryRepository() (user.Repository, error) {
	repo := &userRepository{}
	return repo, nil
}

func (r *userRepository) Create(user *user.User) error {
	return nil
}

func (r *userRepository) GetByEmail(email string) (*user.User, error) {
	us := user.User{ ID: "123456", Name: "Paolo", Email: "paolo@paolo.com", Password: "123456" }
	if email == us.Email {
		return &us, nil
	}
	return nil, nil
}