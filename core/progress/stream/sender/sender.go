package sender

// Sender has responsibility to send message from different type of stream
type Sender interface {
	Send(string) error
}
