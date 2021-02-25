package api

import (
	servicefactories "bands-api/injection/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type health struct {
	IsOk	bool
	Message	string
}

// RegisterHealthCheckHandler returns a handler struct that handles the api requests for the User Domain
func RegisterHealthCheckHandler(router *chi.Mux) {
	router.Get("/api/is-alive", get)
}

func get(w http.ResponseWriter, r *http.Request) {
	health := health{ IsOk: true, Message: "OK" }
	cod := http.StatusOK
	_, err := servicefactories.CreateUserService()	
	if err != nil {
		cod = http.StatusInternalServerError
		health.IsOk = false
		health.Message = err.Error()
	}
	res, _ := json.Marshal(health)
	w.WriteHeader(cod)
	w.Write(res)
}
