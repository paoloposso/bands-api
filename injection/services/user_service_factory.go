package servicefactories

import (
	"bands-api/domain/user"
)

// CreateUserService is a Factory Method that returns a User Service implementation
func CreateUserService(repo user.Repository) (user.Service, error) {
	return user.NewUserService(repo), nil
}

// func chooseRepo() (user.Repository, error) {
// 	env := os.Getenv("ENV")
// 	if env == "TEST" {
// 		return repositorymemory.NewMemoryRepository()
// 	}
// 	mongoURL := os.Getenv("MONGO_URL")
// 	mongoDB := os.Getenv("MONGO_DB")
// 	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
// 	return repositorymongodb.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
// }