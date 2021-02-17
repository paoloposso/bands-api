package test

import (
	"fmt"
	"testing"

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

	fmt.Println(user.ID)

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