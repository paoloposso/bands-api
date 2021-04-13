package test

import (
	"os"
	"strings"
	"testing"

	"github.com/paoloposso/bands-auth-api/user"

	repo "github.com/paoloposso/bands-auth-api/infrastructure/repository/memory"
	"github.com/paoloposso/bands-auth-api/infrastructure/tokenization"

	"github.com/golobby/container"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load("../.env.test")
var service user.Service

func inject(){
	container.Singleton(func() user.Repository {
		repo, err := repo.NewMemoryRepository()
		if err != nil {
			panic(err)
		}
		return repo
	})
	container.Transient(func() user.TokenizationService {
		service := tokenization.NewUserLoginTokenizationService()
		return service
	})
	container.Singleton(func() user.Service {
		var repo user.Repository 
		container.Make(&repo)
		var tokenizationService user.TokenizationService 
		container.Make(&tokenizationService)
		service := user.NewUserService(repo, tokenizationService)
		return service
	})

	container.Make(&service)
}

func TestMain(m *testing.M) {
	inject()
	code := m.Run()
	os.Exit(code)
}

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
	_, token, err := service.Login("pvictorsys@gmail.com", "123456")
	loginToken = token
	if token == "" || err != nil {
		t.Fatal(err)
		t.Fail()
	}
}

func Test_ShouldFailLogin(t *testing.T) {
	_, token, err := service.Login("pvictorsys@gmail.com", "12345")
	if token != "" || err == nil {
		t.Fatal("should return error")
		t.Fail()
	}
	if (!strings.Contains(err.Error(), "Inexistent")) {
		t.Fatalf("should return Inexistent, returned %s instead", err.Error())
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

func Test_ShouldReturnErrorDuplicateEmail(t *testing.T) {
	user := user.User{}
	user.Name = "Paolo"
	user.Password = "12345678"
	user.Email = "pvictorsys@gmail.com"

	err := service.Register(&user)
	if err == nil {
		t.Fail()
	}
}