package user

type TokenizationService interface {
	CreateUserToken(email string, id string) (string, error)
	GetUserIDByToken(tokenString string) (string, error)
}