package event

import "strings"

type asyncEvent struct {
	raw string
}

func NewAsyncDeployment(deploymentID string) StreamEvent {
	return &asyncEvent{
		raw: "async: deployment: " + deploymentID,
	}
}

func (e *asyncEvent) String() string {
	return e.raw
}

func (e *asyncEvent) Type() string {
	ss := strings.Split(e.raw, ":")
	return ss[0]
}

func (e *asyncEvent) Status() string {
	ss := strings.Split(e.raw, ":")
	return ss[1]
}

func (e *asyncEvent) Message() string {
	ss := strings.Split(e.raw, ":")
	return strings.Join(ss[2:], ":")
}
