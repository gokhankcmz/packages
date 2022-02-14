package ErrorTypes

import (
	"fmt"
	"net/http"
)

const ErrorFrom = "PackagesApi"

type Error struct {
	PublicError PublicError
	StatusCode  int
	DetailCode  int
	Args        interface{}
}

type PublicError struct {
	From        string
	Description string
	Detail      string
}

func New(from string, statusCode int, detailCode int, description string, args interface{}) *Error {
	return &Error{
		PublicError: PublicError{
			From:        from,
			Description: description,
		},
		StatusCode: statusCode,
		DetailCode: detailCode,
		Args:       args,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("From: %s, Description: %s, DetailCode: %d,  Args: %v", e.PublicError.From, e.PublicError.Description, e.DetailCode, e.Args)
}

func (e *Error) SetArgs(args interface{}) *Error {
	return &Error{
		PublicError: PublicError{
			From:        e.PublicError.From,
			Description: e.PublicError.Description,
			Detail:      e.PublicError.Detail,
		},
		StatusCode: e.StatusCode,
		DetailCode: e.DetailCode,
		Args:       args,
	}
}

func (e *Error) SetDetailCode(detailCode int) *Error {
	return &Error{
		PublicError: PublicError{
			From:        e.PublicError.From,
			Description: e.PublicError.Description,
			Detail:      e.PublicError.Detail,
		},
		StatusCode: e.StatusCode,
		DetailCode: detailCode,
		Args:       e.Args,
	}
}
func (e *Error) SetPublicDetail(detail string) *Error {
	return &Error{
		PublicError: PublicError{
			From:        e.PublicError.From,
			Description: e.PublicError.Description,
			Detail:      detail,
		},
		StatusCode: e.StatusCode,
		DetailCode: e.DetailCode,
		Args:       e.Args,
	}
}

var (
	UnknownError        = New(ErrorFrom, http.StatusInternalServerError, 0, "An unknown error occured. ", nil)
	EntityNotFound      = New(ErrorFrom, http.StatusNoContent, 0, "Entity not found. ", nil)
	InvalidModel        = New(ErrorFrom, http.StatusBadRequest, 0, "Invalid model. ", nil)
	InvalidToken        = New(ErrorFrom, http.StatusUnauthorized, 0, "Invalid Token.", nil)
	InvalidCredentials  = New(ErrorFrom, http.StatusUnauthorized, 0, "Invalid Token.", nil)
	UnauthorizedRequest = New(ErrorFrom, http.StatusForbidden, 0, "You are not authorized to perform this query.", nil)
	CustomerNotFound    = New(ErrorFrom, http.StatusNotFound, 0, "Customer not found.", nil)
	RedirectionFailed   = New(ErrorFrom, http.StatusServiceUnavailable, 0, "Redirection failed", nil)
)
