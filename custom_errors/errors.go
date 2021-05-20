package customerrors

import "fmt"

type ErrorType int

const (
	InvalidTokenError      = 0
	DBConnectionError      = 1
	UnauthorizedError      = 2
	EmailAlreadyTakenError = 3
	InvalidDataError       = 4
)

// DomainError represents a domain business rule error
type DomainError struct {
	Message   string
	ErrorType ErrorType
}

func (e *DomainError) Error() string {
	return fmt.Sprint(e.Message)
}
