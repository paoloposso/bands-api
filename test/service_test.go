package test

import (
	"fmt"
	"testing"

	customerrors "bands-api/custom_errors"
	"bands-api/repository/memory"
	"bands-api/user"
)

func Test_ShouldGenerateUserID(t *testing.T) {
	repo, err := memory.NewMemoryRepository()

	if err != nil {
		fmt.Println(err)
		panic("MemoryRepository could not be injected")
	}
	
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
	repo, err := memory.NewMemoryRepository()

	if err != nil {
		fmt.Println(err)
		panic("MemoryRepository could not be injected")
	}
	
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
	repo, err := memory.NewMemoryRepository()

	if err != nil {
		fmt.Println(err)
		panic("MemoryRepository could not be injected")
	}
	
	service := user.NewUserService(repo)

	token, err := service.Login("paolo@paolo.com", "123456")

	loginToken = token
	
	if token == "" || err != nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldFailLogin(t *testing.T) {
	repo, err := memory.NewMemoryRepository()
	if err != nil {
		fmt.Println(err)
		panic("MemoryRepository could not be injected")
	}
	service := user.NewUserService(repo)
	token, err := service.Login("paolo@paolo.com", "12345asd")
	if token != "" || err == nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldReceiveExpiredTokenError(t *testing.T) {
	repo, err := memory.NewMemoryRepository()

	if err != nil {
		fmt.Println(err)
		panic("MemoryRepository could not be injected")
	}
	service := user.NewUserService(repo)
	token, err := service.CheckLoginWithToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InBhb2xvQHBhb2xvLmNvbSIsImV4cCI6MTYxMzYwNDk0NSwidXNlcl9pZCI6IjEyMzQ1NiIsInVzZXJuYW1lIjoiUGFvbG8ifQ.z2J6ROmJO5a8zFGXPNTK9UeAaktLzhF5Vv8PvRxrDQk")
	fmt.Println(token)
	ok := false
	if err != nil {
		switch err.(type) {
		case *customerrors.TokenExpiredError:
			ok = true
		}
	}
	if !ok {
		t.Error("should have expired token")
	}
}

func Test_ShouldValidateTokenOk(t *testing.T) {
	repo, err := memory.NewMemoryRepository()

	if err != nil {
		fmt.Println(err)
		panic("MemoryRepository could not be injected")
	}
	service := user.NewUserService(repo)
	token, err := service.CheckLoginWithToken(loginToken)
	if err != nil {
		t.Error(err)
	}
	if token == nil {
		t.Error("Should have returned valid token")
	}
}