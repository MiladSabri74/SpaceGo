package repository

import (
	"errors"
	"fmt"
	"sync"

	"GoRunway/models"
	"GoRunway/task"
)

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrDuplicateTaskID = errors.New("task with this ID already exists")
	ErrNilTask         = errors.New("cannot save nil task")
)

type TaskRepository struct {
	// TODO: Define repository fields
	tasks map[string]*task.Task
	mu    sync.RWMutex
}

func NewTaskRepository() *TaskRepository {
	// TODO: Initialize and return a new TaskRepository
	return &TaskRepository{tasks: make(map[string]*task.Task)}
}

func (r *TaskRepository) Save(t *task.Task) error {
	// TODO: Save task or return ErrNilTask/ErrDuplicateTaskID
	if t == nil {
		return ErrNilTask
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[t.ID()]; exists {
		return fmt.Errorf("%w: %s", ErrDuplicateTaskID, t.ID())
	}

	r.tasks[t.ID()] = t
	return nil
}

func (r *TaskRepository) FindByID(id string) (*task.Task, error) {
	// TODO: Find task by ID or return ErrTaskNotFound
	r.mu.RLock()
	defer r.mu.RUnlock()

	if t, exists := r.tasks[id]; exists {
		return t, nil
	}
	return nil, fmt.Errorf("%w: %s", ErrTaskNotFound, id)
}

func (r *TaskRepository) FindAll() []*task.Task {
	// TODO: Return all stored tasks
	r.mu.RLock()
	defer r.mu.RUnlock()
	var tasks []*task.Task

	for _, t := range r.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (r *TaskRepository) FindByFlightID(flightID string) []*task.Task {
	// TODO: Filter and return tasks by flight ID
	r.mu.RLock()
	defer r.mu.RUnlock()
	var tasks []*task.Task

	for _, t := range r.tasks {
		if t.FlightID() == flightID {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (r *TaskRepository) FindByStatus(status models.TaskStatus) []*task.Task {
	// TODO: Filter and return tasks by status
	r.mu.RLock()
	defer r.mu.RUnlock()
	var tasks []*task.Task

	for _, t := range r.tasks {
		if t.Status() == status {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (r *TaskRepository) FindByPriority(priority models.Priority) []*task.Task {
	// TODO: Filter and return tasks by priority
	r.mu.RLock()
	defer r.mu.RUnlock()
	var tasks []*task.Task

	for _, t := range r.tasks {
		if t.Priority() == priority {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (r *TaskRepository) Delete(id string) error {
	// TODO: Delete task by ID or return ErrTaskNotFound
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; exists {
		delete(r.tasks, id)
		return nil
	}

	return fmt.Errorf("%w", ErrTaskNotFound)
}

func (r *TaskRepository) Count() int {
	// TODO: Return total number of tasks
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.tasks)
}

func (r *TaskRepository) CountByStatus(status models.TaskStatus) int {
	// TODO: Return count of tasks with matching status
	return len(r.FindByStatus(status))
}
