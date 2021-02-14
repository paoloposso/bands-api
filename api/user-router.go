package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/paoloposso/bands-api/user"
)

type UserHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	userService user.UserService
}

func NewHandler(userService user.UserService) UserHandler {
	return &handler {userService: userService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statuscode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statuscode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
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