package wpool

import "sync"

type WorkerPool struct {
	wg sync.WaitGroup
	mu sync.Mutex

	taskQueue  chan func()
	maxWorkers int
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers: maxWorkers,
		taskQueue:  make(chan func()),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.maxWorkers; i++ {
		go func() {
			for task := range wp.taskQueue {
				task()
				wp.wg.Done()
			}
		}()
	}
}

func (wp *WorkerPool) Submit(task func()) {
	wp.wg.Add(1)
	wp.taskQueue <- task
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func (wp *WorkerPool) Stop() {
	close(wp.taskQueue)
}
