package domain

import "strconv"

type DomainError struct {
	Code domainErrorCode
}

type domainErrorCode int

const (
	UnExpectedError domainErrorCode = iota

	// UserError
	UserNotExists
	UsernameIsExists
	UsernameOrPasswordWrong

	// TaskError
	TaskNotExists

	// TaskError
	TagNotExists
)

func NewDomainError(err domainErrorCode) error {
	return &DomainError{
		Code: err,
	}
}

func (e *DomainError) Error() string {
	return "The domain error code: " + strconv.Itoa(int(e.Code))
}
