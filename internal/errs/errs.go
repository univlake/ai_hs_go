package errs

import (
	"errors"
	"fmt"
)

type Code string

const (
	CodeBadRequest      Code = "bad_request"
	CodeUnauthorized    Code = "unauthorized"
	CodeForbidden       Code = "forbidden"
	CodeNotFound        Code = "not_found"
	CodeConflict        Code = "conflict"
	CodeUnprocessable   Code = "unprocessable"
	CodeInternal        Code = "internal"
	CodeCompanyDisabled Code = "company_disabled"
	CodeCrossTenant     Code = "cross_tenant"
	CodeMustChangeCreds Code = "must_change_credentials"
	CodeUsernameTaken   Code = "username_taken"
	CodeUsernameInvalid Code = "username_invalid"
	CodeUsernameReserved Code = "username_reserved"
)

type AppError struct {
	Code    Code
	Message string
	Reason  string
	Cause   error
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error { return e.Cause }

func New(code Code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Wrap(code Code, message string, cause error) *AppError {
	return &AppError{Code: code, Message: message, Cause: cause}
}

func WithReason(code Code, message, reason string) *AppError {
	return &AppError{Code: code, Message: message, Reason: reason}
}

func As(err error) (*AppError, bool) {
	var ae *AppError
	if errors.As(err, &ae) {
		return ae, true
	}
	return nil, false
}
