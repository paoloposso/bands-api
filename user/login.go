package user

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createToken(user User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	atClaims["username"] = user.Name
	atClaims["user_id"] = user.ID
	atClaims["email"] = user.Email
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
	   return nil, err
	}
	return token, nil
}