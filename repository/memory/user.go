package repositorymemory

import (
	"bands-api/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type userMemoryRepository struct {}

// NewMemoryRepository returns a reference to an implementation of UserRepository interface that constains mocked informations, used for tests
func NewMemoryRepository() (user.Repository, error) {
	repo := &userMemoryRepository{}
	return repo, nil
}

func (r *userMemoryRepository) Create(user *user.User) error {
	return nil
}

func (r *userMemoryRepository) GetByEmail(email string) (*user.User, error) {
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