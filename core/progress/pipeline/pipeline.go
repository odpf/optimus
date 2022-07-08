package pipeline

import (
	"github.com/odpf/salt/log"

	"github.com/odpf/optimus/core/progress"
)

type pipelineLogObserver struct {
	log log.Logger
}

func New(log log.Logger) progress.Observer {
	return &pipelineLogObserver{
		log: log,
	}
}

func (obs *pipelineLogObserver) Notify(evt progress.Event) {
	obs.log.Info("observing pipeline log", "progress event", evt.String(), "reporter", "pipeline")
}
