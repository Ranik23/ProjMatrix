package worker

import (
	"ProjMatrix/internal/entity"
	pb "ProjMatrix/pkg/proto"
	"ProjMatrix/pkg/wpool"
)

type Service struct {
	id     string
	status entity.WorkerStatus
	pool   *wpool.WorkerPool
	pb.UnimplementedWorkerServiceServer
}

func NewWorkerService(id string) *Service {
	return &Service{
		id:     id,
		status: entity.WorkerStatusReady,
		pool:   wpool.NewWorkerPool(1),
	}
}

func (s *Service) GetId() string {
	return s.id

}
