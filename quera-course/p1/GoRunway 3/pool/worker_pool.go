package pool

import (
	"errors"
	"fmt"
	"sync"

	"GoRunway/task"
)

var (
	ErrPoolStopped        = errors.New("worker pool is stopped")
	ErrPoolAlreadyStarted = errors.New("worker pool is already running")
	ErrQueueFull          = errors.New("task queue is full")
	ErrInvalidWorkerCount = errors.New("worker count must be greater than zero")
	ErrInvalidCapacity    = errors.New("queue capacity must be greater than zero")
	ErrNilTask            = errors.New("cannot submit nil task")
)

type WorkerPool struct {
	// TODO: Define worker pool fields
	taskQueue   chan (*task.Task)
	workerCount int
	isRunning   bool
	mu          sync.RWMutex
	wg          sync.WaitGroup
}

func NewWorkerPool(workerCount, queueCapacity int) (*WorkerPool, error) {
	// TODO: Validate arguments and return a new WorkerPool
	if workerCount <= 0 {
		return nil,
			fmt.Errorf("%w :%d",
				ErrInvalidWorkerCount.Error(),
				workerCount)
	}
	if queueCapacity <= 0 {
		return nil,
			fmt.Errorf("%w :%d",
				ErrInvalidCapacity.Error(),
				queueCapacity)
	}

	return &WorkerPool{
			taskQueue:   make(chan *task.Task, queueCapacity),
			workerCount: workerCount,
			isRunning:   false},
		nil
}

func (p *WorkerPool) Start() error {
	// TODO: Start worker goroutines
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.isRunning {
		return ErrPoolAlreadyStarted
	}

	p.isRunning = true
	for i := 0; i < p.workerCount; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for t := range p.taskQueue {
				t.Execute()
			}
		}()
	}
	return nil
}

func (p *WorkerPool) Submit(t *task.Task) error {
	// TODO: Submit task to queue or return error
	if t == nil {
		return ErrNilTask
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.isRunning {
		return ErrPoolStopped
	}

	select {
	case p.taskQueue <- t:
		return nil
	default:
		return ErrQueueFull
	}
}

func (p *WorkerPool) Stop() error {
	// TODO: Gracefully stop pool and drain queue
	p.mu.Lock()

	if !p.isRunning {
		return ErrPoolStopped
	}
	p.isRunning = false
	close(p.taskQueue)
	p.mu.Unlock()
	p.wg.Wait()
	return nil
}

func (p *WorkerPool) WorkerCount() int {
	// TODO: Return configured worker count
	return p.workerCount
}

func (p *WorkerPool) IsRunning() bool {
	// TODO: Return running state
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.isRunning
}
