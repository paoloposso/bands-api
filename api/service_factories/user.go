package servicefactories

import (
	repositorymemory "bands-api/repository/memory"
	"bands-api/user"
	"os"
)

// CreateUserService is a Factory Method that returns a User Service implementation
func CreateUserService() user.Service {
	repo, err := chooseRepo()
	if err != nil {
		panic(err)
	}
	return user.NewUserService(repo)
}

func chooseRepo() (user.Repository, error) {
	env := os.Getenv("ENV")
	if env == "TEST" {
		return repositorymemory.NewMemoryRepository()
	}
	return repositorymemory.NewMemoryRepository()
}