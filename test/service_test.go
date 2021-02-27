package test

import (
	"strings"
	"testing"

	"bands-api/domain/user"
	"bands-api/domain/user/login"
	servicefactories "bands-api/injection/services"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load("../.env.test")
var service, _ = servicefactories.CreateUserService()

func Test_ShouldGenerateUserID(t *testing.T) {
	
	user := user.User{}
	user.Name = "Paolo"
	user.Password = "12345678"
	user.Email = "paolo@paolo.com"

	service.Register(&user)
	if user.ID == "" {
		t.Fail()
	}
}

func Test_ShouldFailUserValidation(t *testing.T) {
	user := user.User{}
	user.Name = "Paolo"
	user.Password = "123456"
	user.Email = "paolo@paolo.com"

	err := service.Register(&user)

	if err == nil {
		t.Fail()
	}
}

var loginToken string = ""
func Test_ShouldPerformLogin(t *testing.T) {
	token, err := service.Login(login.Login{ Email: "pvictorsys@gmail.com", Password: "123456" })
	loginToken = token
	if token == "" || err != nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldFailLogin(t *testing.T) {
	token, err := service.Login(login.Login{ Email: "pvictorsys@gmail.com", Password: "12345678" })
	if token != "" || err == nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldReceiveTokenError(t *testing.T) {
	_, err := service.GetDataByToken(strings.Replace(strings.Replace(loginToken, "a", "b", 2), "1", "x", 2))
	if err == nil {
		t.Error("should not validate token")
	}
}

func Test_ShouldValidateTokenOk(t *testing.T) {
	if user, err := service.GetDataByToken(loginToken); err != nil || user == nil {
		t.Error(err)
	}
}