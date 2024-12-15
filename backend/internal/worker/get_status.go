package worker

import (
	"ProjMatrix/internal/entity"
	desc "ProjMatrix/pkg/proto"
	"context"
)

func (s *Service) GetStatus(ctx context.Context, req *desc.GetStatusRequest) (*desc.GetStatusResponse, error) {
	return &desc.GetStatusResponse{
		WorkerId: s.GetId(),
		Status:   string(entity.WorkerStatusReady),
	}, nil
}
