package test

import (
	"fmt"
	"testing"

	"bands-api/repository/memory"
	"bands-api/user"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load("../.env.test")
var repo, err = memory.NewMemoryRepository()

func Test_ShouldGenerateUserID(t *testing.T) {
	service := user.NewUserService(repo)
	
	user := user.User{}
	user.Name = "Paolo"
	user.Password = "123456"
	user.Email = "paolo@paolo.com"

	service.Register(&user)
	if user.ID == "" {
		t.Fail()
	}
}

func Test_ShouldFailUserValidation(t *testing.T) {
	service := user.NewUserService(repo)
	
	user := user.User{}
	user.Name = "Paolo"
	user.Password = "123456"
	user.Email = "paolo@paolo.com"

	err = service.Register(&user)

	if err == nil {
		t.Fail()
	}
}

var loginToken string = ""
func Test_ShouldPerformLogin(t *testing.T) {
	service := user.NewUserService(repo)
	token, err := service.Login("paolo@paolo.com", "123456")
	loginToken = token
	fmt.Println(loginToken)
	if token == "" || err != nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldFailLogin(t *testing.T) {
	service := user.NewUserService(repo)
	token, err := service.Login("paolo@paolo.com", "12345asd")
	if token != "" || err == nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldReceiveTokenError(t *testing.T) {
	service := user.NewUserService(repo)
	_, err = service.CheckLoginWithToken(fmt.Sprintf(loginToken))
	if err == nil {
		t.Error("should have expired token")
	}
}

func Test_ShouldValidateTokenOk(t *testing.T) {
	service := user.NewUserService(repo)
	token, err := service.CheckLoginWithToken(loginToken)
	if err != nil {
		t.Error(err)
	}
	if token == nil {
		t.Error("Should have returned valid token")
	}
}