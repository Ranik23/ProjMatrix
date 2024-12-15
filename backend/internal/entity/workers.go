package entity

import "ProjMatrix/pkg/proto"

type Worker struct {
	Client    proto.WorkerServiceClient // Протовский воркер
	Valuation int64                     // Загруженность воркера
}

type WorkersClient struct {
	FirstWorker  Worker // Первый воркер
	SecondWorker Worker // Второй воркер
}

func (wc *WorkersClient) GetLeastLoadedWorker() Worker {
	if wc.FirstWorker.Valuation <= wc.SecondWorker.Valuation {
		return wc.FirstWorker
	}
	return wc.SecondWorker
}
