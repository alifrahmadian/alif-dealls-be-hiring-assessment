package errors

import "errors"

var (
	ErrNoAuthorizationHeader = errors.New("authorization header is required")
	ErrInvalidTokenFormat    = errors.New("invalid token format")
	ErrTokenInvalid          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token has expired")
	ErrForbidden             = errors.New("access denied")
)
