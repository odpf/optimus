package event

import (
	"fmt"
)

type eventError struct {
	err error
}

func NewError(err error) StreamEvent {
	return &eventError{err}
}

func (e *eventError) String() string {
	return fmt.Sprintf("error: %s", e.err.Error())
}

func (e *eventError) Status() string {
	return "error"
}

func (e *eventError) Message() string {
	return e.err.Error()
}
