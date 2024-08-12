package errorf

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
)

func wrapError(err error, message string, a ...any) error {
	return fmt.Errorf("%s: %w", message, err)
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsNotFoundf(message string, a ...any) error {
	return wrapError(ErrNotFound, message, a...)
}

func IsInternalServer(err error) bool {
	return errors.Is(err, ErrInternalServer)
}

func IsInternalServerf(message string, a ...any) error {
	return wrapError(ErrInternalServer, message, a...)
}

func IsBadRequest(err error) bool {
	return errors.Is(err, ErrBadRequest)
}

func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}
