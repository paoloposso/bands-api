package servicefactories

import (
	customerrors "bands-api/custom_errors"
	repositorymemory "bands-api/repository/memory"
	repositorymongodb "bands-api/repository/mongodb"
	"bands-api/user"
	"os"
	"strconv"
)

// CreateUserService is a Factory Method that returns a User Service implementation
func CreateUserService() user.Service {
	repo, err := chooseRepo()
	if err != nil {
		panic(customerrors.DBConnectionError{ Err: err })
	}
	return user.NewUserService(repo)
}

func chooseRepo() (user.Repository, error) {
	env := os.Getenv("ENV")
	if env == "TEST" {
		return repositorymemory.NewMemoryRepository()
	}
	mongoURL := os.Getenv("MONGO_URL")
	mongoDB := os.Getenv("MONGO_DB")
	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	return repositorymongodb.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
}