package stream

import (
	"sync"

	"github.com/odpf/optimus/core/progress"
	"github.com/odpf/optimus/core/progress/stream/sender"
	"github.com/odpf/salt/log"
)

type streamObserver struct {
	sender sender.Sender
	log    log.Logger
	mu     *sync.Mutex
}

func NewObserver(log log.Logger, sender sender.Sender) progress.Observer {
	return &streamObserver{
		sender: sender,
		log:    log,
		mu:     new(sync.Mutex),
	}
}

func (obs *streamObserver) Notify(evt progress.Event) {
	if err := obs.sender.Send(evt.String()); err != nil {
		obs.log.Error(err.Error())
	}
}
