package api

import (
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
	router.Get("/api/v1/is-alive", get)
}

// health-check - Health Check
// @Summary This API can be used as health check for this application.
// @Description Tells if the auth APIs are working or not.
// @Tags Health Check
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router /is-alive [get]
func get(w http.ResponseWriter, r *http.Request) {
	health := health{ IsOk: true, Message: "OK" }
	cod := http.StatusOK
	// var service user.Service
	// _, err := container.Make(&service)
	// if err != nil {
	// 	cod = http.StatusInternalServerError
	// 	health.IsOk = false
	// 	health.Message = err.Error()
	// }
	res, _ := json.Marshal(health)
	w.WriteHeader(cod)
	w.Write(res)
}
