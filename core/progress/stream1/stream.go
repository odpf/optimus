package stream

import (
	"sync"

	"github.com/odpf/optimus/core/progress"
	pb "github.com/odpf/optimus/core/progress/stream1/protoexample"
	"github.com/odpf/salt/log"
)

type streamObserver struct {
	stream pb.ExampleService_exampleClient
	log    log.Logger
	mu     *sync.Mutex
}

func New(log log.Logger, stream pb.ExampleService_exampleClient) progress.Observer {
	return &streamObserver{
		stream: stream,
		log:    log,
		mu:     new(sync.Mutex),
	}
}

func (obs *streamObserver) Notify(evt progress.Event) {
	if err := obs.stream.Send(constructResponse(evt)); err != nil {
		obs.log.Error(err.Error())
	}
}

func constructResponse(evt progress.Event) pb.ExampleResponse {
	return pb.ExampleResponse{
		Msg: evt.String(),
	}
}
