package api

// LoginRequest has the Login information provided by the requesting User
type LoginRequest struct {
	Email 		string `validate:"required|email"`
	Password 	string `validate:"required|minLen:8"`
}