package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	api "bands-api/api"
	"bands-api/repository/memory"
	"bands-api/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	repo, _ := chooseRepo()
	api.RegisterUserHandler(user.NewUserService(repo), router)
	
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

func chooseRepo() (user.Repository, error) {
	env := os.Getenv("ENV")
	fmt.Println(env)
	if env == "TEST" {
		return memory.NewMemoryRepository()
	}
	return memory.NewMemoryRepository()
}