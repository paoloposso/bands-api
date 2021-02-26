package login

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type loginService struct {}

// CreateToken creates a token for the logged User
func CreateToken(email string, id string) (string, error) {
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

// VerifyToken gets a token sent by request and verifies if it's valid
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
	   return nil, err
	}
	return token, nil
}
