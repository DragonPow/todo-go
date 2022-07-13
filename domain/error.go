package domain

import (
	"errors"
)

var (
	ErrUnexpected = errors.New("ErrUnexpected")

	// Token
	ErrTokenRequired      = errors.New("ErrTokenRequired")
	ErrTokenFormatInvalid = errors.New("ErrTokenFormatInvalid")
	ErrTokenInvalid       = errors.New("ErrTokenInvalid")

	// UserError
	ErrUserNotExists           = errors.New("ErrUserNotExists")
	ErrUserIsExists            = errors.New("ErrUserIsExists")
	ErrUsernameOrPasswordWrong = errors.New("ErrUsernameOrPasswordWrong")

	// TaskError
	ErrTaskNotExists = errors.New("ErrTaskNotExists")

	// TaskError
	ErrTagNotExists       = errors.New("ErrTagNotExists")
	ErrTagValueDuplicated = errors.New("ErrTagValueDuplicated")

	// Association
	ErrTagStillReference = errors.New("ErrTagStillReference")
)
