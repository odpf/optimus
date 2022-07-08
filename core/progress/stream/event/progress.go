package event

import "strings"

type progressEvent struct {
	raw string
}

func NewProgressError(err error) StreamEvent {
	return &progressEvent{
		raw: "progress: error: " + err.Error(),
	}
}

func NewProgressSuccess(msg string) StreamEvent {
	return &progressEvent{
		raw: "progress: success: " + msg,
	}
}

func NewProgressWarn(msg string) StreamEvent {
	return &progressEvent{
		raw: "progress: warning: " + msg,
	}
}

func (e *progressEvent) String() string {
	return e.raw
}

func (e *progressEvent) Type() string {
	ss := strings.Split(e.raw, ":")
	return ss[0]
}

func (e *progressEvent) Status() string {
	ss := strings.Split(e.raw, ":")
	return ss[1]
}

func (e *progressEvent) Message() string {
	ss := strings.Split(e.raw, ":")
	return strings.Join(ss[2:], ":")
}
