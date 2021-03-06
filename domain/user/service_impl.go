package user

import (
	customerrors "bands-api/custom_errors"
	"bands-api/domain/user/login"

	"strings"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo Repository
}

// NewUserService returns a reference to userService struct
func NewUserService(userRepo Repository) Service {
	return &userService{
		userRepo,
	}
}

// Register method performs User's creation in the system
func (s *userService) Register(user *User) error {
	err := user.ValidateRegister()
	if err != nil {
		return err
	}
	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return err
	}
	user.ID = generateID()
	user.Password = hash
	err = s.userRepo.Create(user)
	user.Password = ""
	return err
}

// Login method performs System User's Login, receiving the credentials and returning the token
func (s *userService) Login(loginData login.Login) (string, error) {
	err := loginData.ValidateLogin()
	if err != nil {
		return "", err
	}
	user, err := s.userRepo.GetByEmail(loginData.Email)
	if err != nil {
		return "", err
	} else if user == nil || user.Email == "" {
		return "", &customerrors.UnauthorizedError { Err: errors.New("Inexistent e-mail or wrong password") }
	}
	if checkPasswordHash(loginData.Password, user.Password) {
		token, err := login.CreateToken(user.Email, user.ID)
		if err != nil {
			return "", errors.New("Error creating Token :" + err.Error())
		}
		return token, nil
	} else {
		return "", &customerrors.UnauthorizedError { Err: errors.New("Inexistent e-mail or wrong password") }
	}
}

// GetDataByToken gets the User data (except password) by receiving the authorization token obtained via Login
func (s *userService) GetDataByToken(token string) (*User, error) {
	id, err := login.GetIDByToken(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	user, err := s.userRepo.GetByID(id)
	return user, nil
}

func generateID() string {
    return strings.Replace(uuid.New().String(), "-", "", -1)
}

func generatePasswordHash(plainTextPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 10)
	return string(bytes), err
}

func checkPasswordHash(plaintTextPassword, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintTextPassword))
    return err == nil
}
