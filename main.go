package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	api "bands-api/api"
	factories "bands-api/injection/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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

	service, err := factories.CreateUserService()
	if err != nil {
		panic(err)
	}
	api.RegisterUserHandler(service, router)
	api.RegisterHealthCheckHandler(router)

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
