package domain

import (
	"errors"
)

var (
	ErrUnexpected = errors.New("ErrUnexpected")

	// UserError
	ErrUserNotExists           = errors.New("ErrUserNotExists")
	ErrUserIsExists            = errors.New("ErrUserIsExists")
	ErrUsernameOrPasswordWrong = errors.New("ErrUsernameOrPasswordWrong")

	// TaskError
	ErrTaskNotExists = errors.New("ErrTaskNotExists")

	// TaskError
	ErrTagNotExists = errors.New("ErrTagNotExists")
)
