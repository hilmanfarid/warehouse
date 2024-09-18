package app_errors

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ErrorDetails struct {
	Message     string
	Code        string
	OrinalError string
}

type BaseError struct {
	detail     *ErrorDetails
	stackTrace string
	error
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewBaseError(detail *ErrorDetails) *BaseError {
	berr := &BaseError{detail: detail}
	err := errors.New(detail.Message)
	stracer, ok := err.(stackTracer)
	var sb strings.Builder
	if ok {
		for k := range stracer.StackTrace() {
			if k < 2 {
				continue
			}
			v := stracer.StackTrace()[k] - 1
			f := runtime.FuncForPC(uintptr(v))
			file, line := f.FileLine(uintptr(v))
			if strings.Contains(file, "/go/pkg/") {
				break
			}
			sb.WriteString(f.Name())
			sb.WriteString(" ")
			sb.WriteString(file)
			sb.WriteString(":")
			sb.WriteString(strconv.Itoa(line))
			sb.WriteString("\n")
		}

		berr.stackTrace = sb.String()
	}
	berr.error = detail
	return berr
}

func (b *BaseError) StackTrace() string {
	return b.stackTrace
}

func (b *BaseError) Detail() *ErrorDetails {
	return b.detail
}

func NewErrorDetails(message, originalError, code string) *ErrorDetails {
	return &ErrorDetails{
		OrinalError: originalError,
		Message:     message,
		Code:        code,
	}
}

func NewGeneralError(err error) *BaseError {
	detail := &ErrorDetails{
		OrinalError: err.Error(),
		Message:     "General Error",
		Code:        "422",
	}
	return NewBaseError(detail)
}

func (e *ErrorDetails) Error() string {
	return e.OrinalError
}

var (
	GeneralError              = NewErrorDetails("general error", "", "422")
	AuthorizationInvalidError = NewAuthorizationInvalid(errors.New("Invalid email and password combination"))
)

func NewAuthorizationInvalid(err error) error {
	errDetails := NewErrorDetails("unauthorized", err.Error(), "401")
	return NewBaseError(errDetails)
}

func NewInvalidStateError(err error) error {
	if err == nil {
		err = errors.New("invalid state")
	}
	errDetails := NewErrorDetails("invalid state", err.Error(), "404")
	return NewBaseError(errDetails)
}

func NewInvalidEmailError(err error) error {
	if err == nil {
		err = errors.New("invalid email")
	}
	errDetails := NewErrorDetails("invalid email", err.Error(), "422")
	return NewBaseError(errDetails)
}
