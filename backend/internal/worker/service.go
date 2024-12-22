package worker

import (
	"ProjMatrix/internal/entity"
	pb "ProjMatrix/pkg/proto"
	"ProjMatrix/pkg/repository"
	"ProjMatrix/pkg/wpool"
)

type Service struct {
	id     string
	status entity.WorkerStatus
	Wp     *wpool.WorkerPool
	pb.UnimplementedWorkerServiceServer
	PgRepository repository.PgRepository
}

func NewWorkerService(id string, pgRepository repository.PgRepository) *Service {
	return &Service{
		id:           id,
		status:       entity.WorkerStatusReady,
		Wp:           wpool.NewWorkerPool(8),
		PgRepository: pgRepository,
	}
}

func (s *Service) GetId() string {
	return s.id

}
