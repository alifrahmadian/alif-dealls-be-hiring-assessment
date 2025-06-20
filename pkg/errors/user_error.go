package errors

import "errors"

var (
	ErrUsernameRequired = errors.New("username is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)