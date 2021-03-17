package api

// LoginResponse is a response object that delivers the logged information to the requesting User
type ValidateTokenResponse struct {
	Id	string	`json:"id"`
	Name 	string	`json:"name"`
	Email 	string	`json:"email"`
}