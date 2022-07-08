package event

import (
	"errors"
	"strings"

	"github.com/odpf/optimus/core/progress"
)

type StreamEvent interface {
	progress.Event
	Status() string
	Message() string
}

// Adapter to convert string to StreamEvent
// error: example message
// success: example message
// warning: example message
func GetEvent(s string) StreamEvent {
	switch ss := strings.Split(s, ":"); ss[0] {
	case "error":
		return NewError(errors.New(ss[1]))
	case "success":
		return NewSuccess(ss[1])
	case "warning":
		return NewWarn(ss[1])
	default:
		return nil
	}
}
