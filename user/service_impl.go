package user

import (
	customerrors "github.com/paoloposso/bands-auth-api/custom_errors"

	"strings"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo Repository
	tokenizationService TokenizationService
}

// NewUserService returns a reference to userService struct
func NewUserService(userRepo Repository, tokenizationService TokenizationService) Service {
	return &userService{
		userRepo,
		tokenizationService,
	}
}

// Register method performs User's creation in the system
func (s *userService) Register(user *User) error {
	err := user.ValidateRegister()
	if err != nil {
		return err
	}
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if (existingUser != nil) {
		return &customerrors.DomainError { Message: "E-mail already taken", ErrorType: customerrors.EmailAlreadyTakenError }
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
func (s *userService) Login(email string, password string) (*User, string, error) {
	err := validateLogin(email, password)
	if err != nil {
		return nil, "", err
	}
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, "", err
	} else if user == nil || user.Email == "" {
		return nil, "", &customerrors.DomainError { 
				Message: "Inexistent e-mail or wrong password",
				ErrorType: customerrors.UnauthorizedError,
			}
	}
	if checkPasswordHash(password, user.Password) {
		token, err := s.tokenizationService.CreateUserToken(user.Email, user.ID)
		if err != nil {
			return nil, "", errors.New("Error creating Token :" + err.Error())
		}

		user.Password = ""
		return user, token, nil
	} else {
		return nil, "", &customerrors.DomainError { 
				Message: "Inexistent e-mail or wrong password",
				ErrorType: customerrors.UnauthorizedError,
			}
	}
}

// GetDataByToken gets the User data (except password) by receiving the authorization token obtained via Login
func (s *userService) GetDataByToken(token string) (*User, error) {
	id, err := s.tokenizationService.GetUserIDByToken(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	user, err := s.userRepo.GetByID(id)

	user.Password = ""
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

func validateLogin(email string, password string) error {
	var errors []string
	
	if email == "" {
		errors = append(errors, "E-mail is required")
	}
	if password == "" {
		errors = append(errors, "Password is required")
	}
	if (len(errors) > 0) {
		return &customerrors.DomainError{ Message: strings.Join(errors, ";"), ErrorType: customerrors.InvalidDataError }
	}

	return nil
}