package memory

import (
	"bands-api/user"

	"golang.org/x/crypto/bcrypt"
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
	hash, _ := generatePasswordHash("123456")
	us := user.User{ ID: "123456", Name: "Paolo", Email: "paolo@paolo.com", Password: hash }
	if email == us.Email {
		return &us, nil
	}
	return nil, nil
}

func generatePasswordHash(plainTextPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 14)
	return string(bytes), err
}