package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/odpf/optimus/store"
)

const (
	ErrNotFound        ErrorType = "Not Found"
	ErrAlreadyExists   ErrorType = "Resource Already Exists"
	ErrInvalidArgument ErrorType = "Invalid Argument"
	ErrInternalError   ErrorType = "Internal Error"
	ErrFailedPrecond   ErrorType = "Failed Precondition"
)

type ErrorType string

func (s ErrorType) String() string {
	return strings.ToLower(string(s))
}

// DomainError is used to map different type of errors identified in service to network errors
type DomainError struct {
	Entity    string
	ErrorType ErrorType
	Message   string
	Err       error
}

func NewError(entity string, errType ErrorType, msg string) *DomainError {
	return &DomainError{
		Entity:    entity,
		ErrorType: errType,
		Message:   msg,
		Err:       nil,
	}
}

func FromError(err error, entity, msg string) *DomainError {
	errType := ErrInternalError
	msgStr := "internal error"
	if errors.Is(err, store.ErrResourceNotFound) {
		errType = ErrNotFound
		msgStr = err.Error()
	} else if errors.Is(err, store.ErrResourceExists) {
		errType = ErrAlreadyExists
		msgStr = err.Error()
	} else if errors.Is(err, store.ErrEmptyConfig) {
		errType = ErrFailedPrecond
		msgStr = err.Error()
	}

	// TODO: Improve the msg, append to existing msg
	if msg == "" {
		msg = msgStr
	}

	return &DomainError{
		Err:       err,
		Message:   msg,
		Entity:    entity,
		ErrorType: errType,
	}
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%v: %v for entity %v",
		e.Message, e.ErrorType.String(), e.Entity)
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

func (e *DomainError) DebugString() string {
	var wrappedError string
	var de *DomainError
	if errors.As(e.Err, &de) {
		wrappedError = de.DebugString()
	} else if e.Err != nil {
		wrappedError = e.Err.Error()
	}

	return fmt.Sprintf("%v: %v for %v: %s",
		e.Message, e.ErrorType.String(), e.Entity, wrappedError)
}
