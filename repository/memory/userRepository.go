package memory

import (
	"bands-api/user"
)

type userRepository struct {}

func NewMemoryRepository() (user.UserRepository, error) {
	repo := &userRepository{}
	return repo, nil
}

func (r *userRepository) Create(user *user.User) error {
	return nil
}