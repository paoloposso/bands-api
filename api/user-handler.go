package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bands-api/api/model"
	"bands-api/user"

	customerrors "bands-api/custom_errors"

	"github.com/go-chi/chi"
)

type userHandler struct {
	userService user.Service
}

// RegisterUserHandler returns a handler struct that handles the api requests for the User Domain
func RegisterUserHandler(userService user.Service, router *chi.Mux) {
	handler := userHandler {userService: userService}
	router.Post("/api/user", handler.create)
	router.Post("/api/user/login", handler.login)
}

func (h *userHandler) create(w http.ResponseWriter, r *http.Request) {
	var u user.User
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

func (h *userHandler) login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	token, err := h.userService.Login(req.Email, req.Password)
	w.WriteHeader(http.StatusOK)
	w.Write(encodeToBytes(model.LoginResponse { Token: token }))
}

func formatError(err error) (int, string) {
	code := http.StatusInternalServerError
	switch err.(type) {
		case *customerrors.InvalidDataError:
			code = http.StatusBadRequest
		case *customerrors.InvalidEmailOrIncorrectPasswordError:
			code = http.StatusForbidden
	}
	return code, fmt.Sprintf("{ message: \"%s\" }", err)
}
