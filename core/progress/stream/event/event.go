package event

import (
	"strings"

	"github.com/odpf/optimus/core/progress"
)

// Medium to convert progress.Event to Event protobuf
type StreamEvent interface { // similar with event protobuf
	progress.Event
	Type() string   // progress, async
	Status() string // success, warning, error
	Message() string
}

// progress: error: example message
// progress: success: example message
// progress: warning: example message
// async: deployment: 12345
func GetEvent(s string) StreamEvent {
	switch ss := strings.Split(s, ":"); ss[0] {
	case "progress":
		// example Success
		return NewProgressSuccess(ss[2])
	case "async":
		// example deployment async
		return NewAsyncDeployment(ss[2])
	}
	return nil
}
