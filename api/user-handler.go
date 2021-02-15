package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/paoloposso/bands-api/user"
)

type UserHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userService user.UserService
}

func RegisterUserHandler(userService user.UserService, router *chi.Mux) {
	handler := userHandler {userService: userService}
	router.Post("/api/user", handler.Post)
}

func (h *userHandler) Post(w http.ResponseWriter, r *http.Request) {
	var u user.User
	fmt.Println(r.Header)
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.userService.Register(&u)
	res, err := json.Marshal(&u)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func formatError(err error) (int, string) {
	return http.StatusInternalServerError, fmt.Sprintf("{ error: %v }", err)
}