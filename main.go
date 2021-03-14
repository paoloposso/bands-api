package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	api "bands-api/api"
	"bands-api/domain/user"

	repositorymongodb "bands-api/infrastructure/repository/mongodb"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/golobby/container"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	injectDependencies(router)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port ", httpPort())
		errs <- http.ListenAndServe(httpPort(), router)
	}()
	go func(){
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func injectDependencies(router *chi.Mux) {
	container.Singleton(func() user.Repository {
		mongoURL := os.Getenv("MONGO_URL")
		mongoDB := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := repositorymongodb.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
		if err != nil {
			panic(err)
		}
		return repo
	})
	container.Singleton(func() user.Service {
		var repo user.Repository 
		container.Make(&repo)
		service := user.NewUserService(repo)
		return service
	})

	var service user.Service
	container.Make(&service)
	
	api.RegisterUserHandler(service, router)
	api.RegisterHealthCheckHandler(router)
}