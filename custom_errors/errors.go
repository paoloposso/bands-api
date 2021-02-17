package customerrors

import "fmt"

// InvalidDataError represents an error thrown when the input data is invalid
type InvalidDataError struct {
    Message string
}

func (e *InvalidDataError) Error() string {
    return fmt.Sprintf("Invalid data: %s", e.Message)
}

// InvalidEmailOrIncorrectPasswordError represents an error thrown when the email is nor registered or password is incorrect
type InvalidEmailOrIncorrectPasswordError struct {
    Email string
}

func (e *InvalidEmailOrIncorrectPasswordError) Error() string {
    return fmt.Sprintf("E-mail %s not found or password is invalid", e.Email)
}