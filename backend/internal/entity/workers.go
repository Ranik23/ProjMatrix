package entity

import (
	"ProjMatrix/pkg/proto"
	"ProjMatrix/pkg/repository"
)

type Worker struct {
	Client    proto.WorkerServiceClient // Протовский воркер
	Valuation int64                     // Загруженность воркера
}

type WorkersClient struct {
	FirstWorker  Worker // Первый воркер
	SecondWorker Worker // Второй воркер
	PgRepository repository.PgRepository
	Session      string
}

func NewWorkersClient(firstWorker proto.WorkerServiceClient, secondWorker proto.WorkerServiceClient, pgRepo repository.PgRepository, session string) *WorkersClient {
	return &WorkersClient{
		FirstWorker: Worker{
			Client:    firstWorker,
			Valuation: 0,
		},
		SecondWorker: Worker{
			Client:    secondWorker,
			Valuation: 0,
		},
		PgRepository: pgRepo,
		Session:      session,
	}
}

func (wc *WorkersClient) GetLeastLoadedWorker() Worker {
	if wc.FirstWorker.Valuation <= wc.SecondWorker.Valuation {
		return wc.FirstWorker
	}
	return wc.SecondWorker
}
