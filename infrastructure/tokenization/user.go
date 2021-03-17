package tokenization

import (
	customerrors "bands-api/custom_errors"
	"bands-api/domain/user"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type userLoginTokenizationService struct {}

func NewUserLoginTokenizationService() user.TokenizationService {
	return &userLoginTokenizationService{}
}

// CreateUserToken creates a token for the logged User
func (u *userLoginTokenizationService) CreateUserToken(email string, id string) (string, error) {
	t, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTES"))
	if err != nil || t == 0 {
		t = 15
	}
	atClaims := jwt.MapClaims{}
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(t)).Unix()
	atClaims["user_id"] = id
	atClaims["email"] = email
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
}

// GetUserIDByToken gets a token sent by request and, if the token is valid, returns the user ID
func (u *userLoginTokenizationService) GetUserIDByToken(tokenString string) (string, error) {
	claims := jwt.MapClaims {}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	   return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
	   return "", err
	}
	if !token.Valid {
		return "", &customerrors.InvalidTokenError {}
	}
	userID := fmt.Sprint(claims["user_id"])
	return userID, nil
}
