package sender

import (
	pb "github.com/odpf/optimus/api/proto/odpf/optimus/core/v1beta1"
)

type resourceStreamSender struct {
	stream pb.ResourceService_DeployResourceSpecificationServer
}

func NewResourceSender(stream pb.ResourceService_DeployResourceSpecificationServer) Sender {
	return &resourceStreamSender{
		stream: stream,
	}
}

func (s *resourceStreamSender) Send(msg string) error {
	resp := &pb.DeployResourceSpecificationResponse{
		
	}
	return s.stream.Send(resp)
}
