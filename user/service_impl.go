package user

import (
	customerrors "bands-api/custom_errors"
	"bands-api/hashing"
	passwordutil "bands-api/password"
)

type userService struct {
	userRepo UserRepository
}

// NewUserService returns a reference to userService struct
func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) Register(user *User) error {
	user.Id = hashing.GenerateUuid()
	hash, err := passwordutil.GeneratePasswordHash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	err = s.userRepo.Create(user)
	user.Password = ""
	return err
}

// Login method performs System User's Login using 
func (s *userService) Login(email string, plainTextPassword string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if passwordutil.CheckPasswordHash(plainTextPassword, user.Password) {
		return "token", nil
	}
	return "", &customerrors.InvalidEmailOrIncorrectPasswordError { Email: email }
}
