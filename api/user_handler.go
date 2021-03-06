package api

import (
	"encoding/json"
	"net/http"

	dto "github.com/paoloposso/bands-auth-api/api/dto"
	"github.com/paoloposso/bands-auth-api/user"

	"github.com/go-chi/chi"
)

type userHandler struct {
	userService user.Service
}

// RegisterUserHandler returns a handler struct that handles the api requests for the User Domain
func RegisterUserHandler(userService user.Service, router *chi.Mux) {
	handler := userHandler {userService: userService}

	baseUrl := "/api/v1"

	router.Post(baseUrl + "/user", handler.register)
	router.Post(baseUrl + "/user/login", handler.login)
	router.Get(baseUrl + "/user/me", handler.validateToken)
}

// user - Registers an User
// @Summary This API can be used to register an User.
// @Description Registers an User.
// @Tags User
// @Accept  json
// @Produce  json
// @Param userRegistration body RegisterRequest true "User Registration"
// @Success 201
// @Router /user [post]
func (h *userHandler) register(w http.ResponseWriter, r *http.Request) {
	var registerReq dto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

	u := user.User {
		Name: registerReq.Name,
		Password: registerReq.Password,
		Email: registerReq.Email,
	}

	err = h.userService.Register(&u)
	
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// validate token - Validates a Token and returns the data
// @Summary This API can be used to validate a token.
// @Description Validate Token
// @Tags User
// @Accept  json
// @Produce  json
// @Param token query string true "Token"
// @Success 200 {object} ValidateTokenResponse "api response"
// @Router /user/me [get]
func (h *userHandler) validateToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	token := r.URL.Query().Get("token")
	user, err := h.userService.GetDataByToken(token)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	dto := dto.ValidateTokenResponse { Id: user.ID, Name: user.Name, Email: user.Email }
	res, err := json.Marshal(dto)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// login - Authenticates an User
// @Summary This API can be used authenticate an User.
// @Description User Login.
// @Tags User
// @Accept  json
// @Produce  json
// @Param login_request body LoginRequest true "Login Request"
// @Success 200 {object} LoginResponse "api response"
// @Router /user/login [post]
func (h *userHandler) login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var login dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	user, token, err := h.userService.Login(login.Email, login.Password)
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	res, err := json.Marshal(&dto.LoginResponse { Token: token, ID: user.ID })
	if err != nil {
		code, msg := formatError(err)
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
