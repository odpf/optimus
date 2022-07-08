package event

import (
	"fmt"
)

type eventWarning struct {
	msg string
}

func NewWarn(msg string) StreamEvent {
	return &eventWarning{msg}
}

func (e *eventWarning) String() string {
	return fmt.Sprintf("warning: %s", e.msg)
}

func (e *eventWarning) Status() string {
	return "warning"
}

func (e *eventWarning) Message() string {
	return e.msg
}
