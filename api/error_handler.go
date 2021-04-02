package api

import (
	customerrors "bands-auth-api/custom_errors"
	"fmt"
	"net/http"
	"reflect"
)

func returnHTTPError(err error) (int, string) {
	domainError := err.(*customerrors.DomainError)
	code := http.StatusInternalServerError 
	switch domainError.ErrorType {
		case customerrors.InvalidDataError:
			code = http.StatusBadRequest
		case customerrors.UnauthorizedError:
			code = http.StatusForbidden
		case customerrors.EmailAlreadyTakenError:
			code = http.StatusConflict
	}
	return code, fmt.Sprintf("{ \"message\": \"%s\" }", err)
}

func formatError(err error) (int, string) {
	code := http.StatusInternalServerError
	errType := reflect.TypeOf(err).String()

	if errType == "*customerrors.DomainError" {
		return returnHTTPError(err)
	}
	
	return code, fmt.Sprintf("{ \"message\": \"%s\" }", err)
}