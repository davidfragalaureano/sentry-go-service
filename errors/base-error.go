package errors

import (
	"fmt"
)

type BaseError struct {
	Err            error
	Message        string
	HttpStatusCode int
	Code           int
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("%d Error: %d Status %v", e.Code, e.HttpStatusCode, e.Message)
}

type UnexpectedError struct {
	BaseError
}

type InvalidSpaceObjectError struct {
	BaseError
}

type IOReadError struct {
	BaseError
}

func NewIOReadError(err error, message string) (e *IOReadError) {
	return &IOReadError{
		BaseError{
			Err:            err,
			Message:        "NewIOReadError: " + message,
			HttpStatusCode: 500,
			Code:           100,
		},
	}
}

func NewUnexpectedError(err error, message string) (e *UnexpectedError) {
	return &UnexpectedError{
		BaseError{
			Err:            err,
			Message:        "UnexpectedError: " + message,
			HttpStatusCode: 500,
			Code:           100,
		},
	}
}

func NewInvalidSpaceObjectError(err error, message string) (e *InvalidSpaceObjectError) {
	return &InvalidSpaceObjectError{
		BaseError{
			Err:            err,
			Message:        "InvalidSpaceObjectError: " + message,
			HttpStatusCode: 400,
			Code:           101,
		},
	}
}
