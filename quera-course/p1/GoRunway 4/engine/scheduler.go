package engine

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"GoRunway/models"
	"GoRunway/pool"
	"GoRunway/repository"
)

var (
	ErrSchedulerStopped   = errors.New("scheduler is stopped")
	ErrInvalidDelay       = errors.New("delay must be positive")
	ErrInvalidInterval    = errors.New("interval must be positive")
	ErrNilPool            = errors.New("pool should not be nil")
	ErrNilRepository      = errors.New("repository should not be nil")
	ErrSchedulerIsRunning = errors.New("scheduler is ruuning now")
)

// FlightReport summarizes the execution statistics of all tasks for a given flight
type FlightReport struct {
	FlightID       string
	TotalTasks     int
	CompletedTasks int
	FailedTasks    int
	CancelledTasks int
	PendingTasks   int
	Errors         []string
}

type Scheduler struct {
	// TODO: Define scheduler fields
	repo      *repository.TaskRepository
	pool      *pool.WorkerPool
	isRunning bool
	stopChan  chan struct{}
	mu        sync.RWMutex
	wg        sync.WaitGroup
}

func NewScheduler(repo *repository.TaskRepository, pool *pool.WorkerPool) (*Scheduler, error) {
	// TODO: Validate arguments and return a new Scheduler
	if repo == nil {
		return nil, ErrNilRepository
	}
	if pool == nil {
		return nil, ErrNilPool
	}
	return &Scheduler{
			repo:      repo,
			pool:      pool,
			stopChan:  make(chan struct{}),
			isRunning: false},
		nil
}

func (s *Scheduler) Start() error {
	// TODO: Start scheduler and ensure worker pool is running
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		return ErrSchedulerIsRunning
	}
	if !s.pool.IsRunning() {
		s.pool.Start()
	}
	s.isRunning = true
	return nil
}

func (s *Scheduler) ScheduleNow(taskID string) error {
	// TODO: Immediately dispatch task to worker pool
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		return ErrSchedulerStopped
	}

	task, err := s.repo.FindByID(taskID)
	if err == nil {
		s.pool.Submit(task)
		return nil
	}
	return fmt.Errorf("find task %q: %w", taskID, err)
}

func (s *Scheduler) ScheduleDelayed(taskID string, delay time.Duration) error {
	// TODO: Schedule task to run after delay
	if delay <= 0 {
		return ErrInvalidDelay
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		return ErrSchedulerStopped
	}

	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return fmt.Errorf("find task %q: %w", taskID, err)
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		taskTimer := time.NewTimer(delay)
		defer taskTimer.Stop()

		select {
		case <-s.stopChan:
			return
		case <-taskTimer.C:
			s.pool.Submit(task)
		}
	}()
	return nil

}

func (s *Scheduler) SchedulePeriodic(taskID string, interval time.Duration, stopRecurring <-chan struct{}) error {
	// TODO: Schedule task to execute periodically until stopRecurring
	if interval <= 0 {
		return ErrInvalidInterval
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		return ErrSchedulerStopped
	}

	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return fmt.Errorf("find task %q: %w", taskID, err)
	}

	s.wg.Add(1)

	go func() {
		defer s.wg.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-s.stopChan:
				return
			case <-stopRecurring:
				return
			case <-ticker.C:
				s.mu.RLock()
				if !s.isRunning {
					s.mu.RUnlock()
					return
				}
				ticker.Reset(interval)
				s.pool.Submit(task)
			}
		}
	}()
	return nil

}

func (s *Scheduler) GetFlightReport(flightID string) FlightReport {
	// TODO: Aggregate and return flight statistics
	tasks := s.repo.FindByFlightID(flightID)
	var failedErrs []string
	var completedTasks, failedTasks, cancelledTasks, pendingTasks int
	for _, t := range tasks {
		if t.Status() == models.StatusCompleted {
			completedTasks++
		} else if t.Status() == models.StatusCancelled {
			cancelledTasks++
		} else if t.Status() == models.StatusFailed {
			failedTasks++
			failedErrs = append(failedErrs, t.Error().Error())
		} else if t.Status() == models.StatusPending {
			pendingTasks++
		} else if t.Status() == models.StatusRunning {
			pendingTasks++
		}
	}
	return FlightReport{
		FlightID:       flightID,
		TotalTasks:     len(tasks),
		CompletedTasks: completedTasks,
		FailedTasks:    failedTasks,
		CancelledTasks: cancelledTasks,
		PendingTasks:   pendingTasks,
		Errors:         failedErrs,
	}
}

func (s *Scheduler) Stop() error {
	// TODO: Gracefully stop scheduler and worker pool
	s.mu.Lock()

	if !s.isRunning {
		s.mu.Unlock()
		return ErrSchedulerStopped
	}
	s.isRunning = false
	close(s.stopChan)
	s.mu.Unlock()

	s.wg.Wait()
	s.pool.Stop()

	return nil
}
