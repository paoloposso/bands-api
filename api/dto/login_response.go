package api

// LoginResponse is a response object that delivers the logged information to the requesting User
type LoginResponse struct {
	Token	string	`json:"token"`
	Email 	string	`json:"email"`
	Name 	string	`json:"name"`
}