package event

import (
	"fmt"
)

type eventSuccess struct {
	msg string
}

func NewSuccess(msg string) StreamEvent {
	return &eventSuccess{msg}
}

func (e *eventSuccess) String() string {
	return fmt.Sprintf("success: %s", e.msg)
}

func (e *eventSuccess) Status() string {
	return "success"
}

func (e *eventSuccess) Message() string {
	return e.msg
}
