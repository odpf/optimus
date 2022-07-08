package sender

import (
	pb "github.com/odpf/optimus/api/proto/odpf/optimus/core/v1beta1"
	"github.com/odpf/optimus/core/progress/stream/event"
)

type jobDeploymentStreamSender struct {
	stream pb.JobSpecificationService_DeployJobSpecificationServer
}

func NewJobDeploymentSender(stream pb.JobSpecificationService_DeployJobSpecificationServer) Sender {
	return &jobDeploymentStreamSender{
		stream: stream,
	}
}

func (s *jobDeploymentStreamSender) Send(msg string) error {
	e := event.GetEvent(msg)
	resp := &pb.DeployJobSpecificationResponse{}
	resp.Event.Status = e.Status()
	resp.Event.Message = e.Message()
	return s.stream.Send(resp)
}
