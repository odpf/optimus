package protoexample

type ExampleResponse struct {
	Msg string
}

type ExampleService_exampleClient interface {
	Send(ExampleResponse) error
}
