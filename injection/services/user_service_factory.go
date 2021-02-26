package servicefactories

import (
	customerrors "bands-api/custom_errors"
	"bands-api/domain/user"
	repositorymemory "bands-api/repository/memory"
	repositorymongodb "bands-api/repository/mongodb"
	"os"
	"strconv"
)

// CreateUserService is a Factory Method that returns a User Service implementation
func CreateUserService() (user.Service, error) {
	repo, err := chooseRepo()
	if err != nil {
		er := &customerrors.DBConnectionError { Err: err }
		return nil, er
	}
	return user.NewUserService(repo), nil
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