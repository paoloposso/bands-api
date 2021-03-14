package customerrors

import "fmt"

// InvalidDataError represents an error thrown when the input data is invalid
type InvalidDataError struct {
    Message string
}
func (e *InvalidDataError) Error() string {
    return fmt.Sprintf("Invalid data: %s", e.Message)
}

// InvalidTokenError represents an error thrown when the email is nor registered or password is incorrect or token is expired
type InvalidTokenError struct {
}
func (e *InvalidTokenError) Error() string {
    return fmt.Sprint("Invalid Token!")
}

// DBConnectionError represents an error thrown when the email is nor registered or password is incorrect
type DBConnectionError struct {
    Err error
}
func (e *DBConnectionError) Error() string {
    return fmt.Sprintf("Error trying to connect to database: %s", e.Err.Error())
}

// UnauthorizedError represents an error thrown when the token is invalid or expired
type UnauthorizedError struct {
    Err error
}
func (e *UnauthorizedError) Error() string {
    return fmt.Sprintf("Unauthorized: %s", e.Err)
}

// EmailAlreadyTakenError represents an error thrown when the informed e-mail already exists
type EmailAlreadyTakenError struct {
    Err error
}
func (e *EmailAlreadyTakenError) Error() string {
    return "E-mail already taken"
}