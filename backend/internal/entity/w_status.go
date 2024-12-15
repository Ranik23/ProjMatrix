package entity

type WorkerStatus string

const (
	WorkerStatusReady   WorkerStatus = "ready"
	WorkerStatusBusy    WorkerStatus = "busy"
	WorkerStatusStopped WorkerStatus = "stopped"
)
