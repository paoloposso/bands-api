package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bands-api/user"

	customerrors "bands-api/custom_errors"

	"github.com/go-chi/chi"
)

// UserHandler constains methods to handle the api requests for the User Domain
type UserHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userService user.Service
}

// RegisterUserHandler returns a handler struct that handles the api requests for the User Domain
func RegisterUserHandler(userService user.Service, router *chi.Mux) {
	handler := userHandler {userService: userService}
	router.Post("/api/user", handler.Post)
}

func (h *userHandler) Post(w http.ResponseWriter, r *http.Request) {
	var u user.User
	fmt.Println(r.Header)
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	err = h.userService.Register(&u)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	res, err := json.Marshal(&u)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func formatError(err error) (int, string) {
	code := http.StatusInternalServerError
	switch err.(type) {
		case *customerrors.InvalidDataError:
			code = http.StatusBadRequest
		case *customerrors.InvalidEmailOrIncorrectPasswordError:
			code = http.StatusNoContent
	}
	return code, fmt.Sprintf("{ message: \"%s\" }", err)
}