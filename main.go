package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	api "bands-auth-api/api"
	"bands-auth-api/infrastructure/tokenization"
	"bands-auth-api/user"

	repositorymongodb "bands-auth-api/infrastructure/repository/mongodb"

	_ "bands-auth-api/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/golobby/container"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Bands Auth API
// @version 1.0
// @description Authentication and Registration API.

// @BasePath /api/v1
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

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition"
	))

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

	fmt.Println("Terminated ", <-errs)
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func injectDependencies(router *chi.Mux) {
	container.Transient(func() user.Repository {
		mongoURL := os.Getenv("MONGO_URL")
		mongoDB := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := repositorymongodb.NewMongoUserRepository(mongoURL, mongoDB, mongoTimeout)
		if err != nil {
			panic(err)
		}
		return repo
	})
	
	container.Transient(func() user.TokenizationService {
		service := tokenization.NewUserLoginTokenizationService()
		return service
	})

	container.Transient(func() user.Service {
		var repo user.Repository 
		container.Make(&repo)
		var tokenizationService user.TokenizationService 
		container.Make(&tokenizationService)
		service := user.NewUserService(repo, tokenizationService)
		return service
	})
	
	var service user.Service
	container.Make(&service)
	
	api.RegisterUserHandler(service, router)
	api.RegisterHealthCheckHandler(router)
}