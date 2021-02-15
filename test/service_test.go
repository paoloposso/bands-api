package test

import (
	"fmt"
	"testing"

	"github.com/paoloposso/bands-api/repository/memory"
	"github.com/paoloposso/bands-api/user"
)

func Test_ShouldGenerateUserId(t *testing.T) {
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

	fmt.Println(user.Id)

	if user.Id == "" {
		t.Fail()
	}
}