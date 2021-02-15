package user

import (
	id "github.com/paoloposso/bands-api/util"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) Register(user *User) error {
	user.Id = id.GenerateUuid()
	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	err = s.userRepo.Create(user)
	user.Password = ""
	return err
}

func generatePasswordHash(plainTextPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}