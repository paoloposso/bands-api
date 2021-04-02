package api

// RegisterRequest has the Login information provided by the requesting User
type RegisterRequest struct {
	Email 		string `validate:"required|email"`
	Name 		string `validate:"required"`
	Password 	string `validate:"required|minLen:8"`
}