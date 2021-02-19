package user

import (
	customerrors "bands-api/custom_errors"
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

func (s *userService) Register(user *User) error {
	user.ID = generateID()
	err := user.ValidateRegister()
	if err != nil {
		return err
	}
	hash, err := generatePasswordHash(user.Password)
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
	if checkPasswordHash(plainTextPassword, user.Password) {
		token, err := createToken(*user)
		if err != nil {
			return "", errors.New("Error creating Token :" + err.Error())
		}
		return token, nil
	}
	return "", &customerrors.InvalidEmailOrIncorrectPasswordError { Email: email }
}

func (s *userService) CheckLoginWithToken(tokenString string) (*jwt.Token, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func generateID() string {
    return strings.Replace(uuid.New().String(), "-", "", -1)
}

func generatePasswordHash(plainTextPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 14)
	return string(bytes), err
}

func checkPasswordHash(plaintTextPassword, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintTextPassword))
    return err == nil
}
