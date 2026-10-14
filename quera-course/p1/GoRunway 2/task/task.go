package task

import (
	"GoRunway/models"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	// TODO: Define struct fields
	id          string
	title       string
	flightID    string
	priority    models.Priority
	status      models.TaskStatus
	action      models.Action
	createdAt   time.Time
	startedAt   time.Time
	completedAt time.Time
	err         error
	mu          sync.RWMutex
}

func NewTask(id, title, flightID string, priority models.Priority, action models.Action) (*Task, error) {
	// TODO: Validate fields and return a new Task in PENDING status or wrapped error
	if id == "" {
		return nil, fmt.Errorf("%w: %w",
			&models.ValidationError{
				Field:  "id",
				Reason: models.ErrInvalidTaskID.Error(),
			},
			models.ErrInvalidTaskID,
		)
	}
	if title == "" {
		return nil, fmt.Errorf("%w: %w",
			&models.ValidationError{
				Field:  "title",
				Reason: models.ErrEmptyTitle.Error(),
			},
			models.ErrEmptyTitle,
		)
	}
	if flightID == "" {
		return nil, fmt.Errorf("%w: %w",
			&models.ValidationError{
				Field:  "flightID",
				Reason: models.ErrEmptyFlightID.Error(),
			},
			models.ErrEmptyFlightID,
		)
	}
	if priority < models.PriorityLow ||
		priority > models.PriorityCritical {
		return nil, fmt.Errorf("%w: %w",
			&models.ValidationError{
				Field:  "priority",
				Reason: models.ErrInvalidPriority.Error(),
			},
			models.ErrInvalidPriority,
		)
	}

	if action == nil {
		return nil, fmt.Errorf("%w: %w",
			&models.ValidationError{
				Field:  "action",
				Reason: models.ErrNilAction.Error(),
			},
			models.ErrNilAction,
		)
	}

	t := &Task{
		id:        id,
		title:     title,
		flightID:  flightID,
		priority:  priority,
		action:    action,
		status:    models.StatusPending,
		createdAt: time.Now(),
	}
	return t, nil
}

func (t *Task) ID() string {
	// TODO: Safely return task ID
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.id
}

func (t *Task) Title() string {
	// TODO: Safely return task title
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.title
}

func (t *Task) FlightID() string {
	// TODO: Safely return flight ID
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.flightID
}

func (t *Task) Priority() models.Priority {
	// TODO: Safely return priority
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.priority
}

func (t *Task) Status() models.TaskStatus {
	// TODO: Safely return status
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.status
}

func (t *Task) Error() error {
	// TODO: Safely return recorded error
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.err
}

func (t *Task) Cancel() error {
	// TODO: Cancel pending task or return error
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.status != models.StatusPending {
		return fmt.Errorf("%w", models.ErrTaskCannotCancel)
	}

	t.status = models.StatusCancelled
	t.completedAt = time.Now()
	return nil
}

func (t *Task) Reset() error {
	// TODO: Reset completed/failed/cancelled task to PENDING or return error
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.status == models.StatusRunning {
		return errors.New("cannot reset running task")
	}

	t.err = nil
	t.status = models.StatusPending
	t.startedAt = time.Time{}
	t.completedAt = time.Time{}

	return nil
}

func (t *Task) Execute() (err error) {
	// TODO: Execute action with panic recovery, update status and timestamps
	t.mu.Lock()

	if t.status != models.StatusPending {
		return models.ErrTaskNotPending
	}

	t.status = models.StatusRunning
	t.startedAt = time.Now()

	action := t.action
	t.mu.Unlock()

	defer func() {
		if r := recover(); r != nil {
			t.status = models.StatusFailed
			t.completedAt = time.Now()
			t.err = fmt.Errorf("task panicked: %v", r)
		}

		t.mu.Lock()
		defer t.mu.Unlock()
		t.completedAt = time.Now()
		if err != nil {
			t.status = models.StatusFailed
			t.err = err
		} else {
			t.status = models.StatusCompleted
			t.err = nil
		}
	}()

	err = action()
	return err
}

func (t *Task) GetInfo() models.TaskInfo {
	// TODO: Return immutable snapshot of task info
	t.mu.RLock()
	defer t.mu.RUnlock()

	info := models.TaskInfo{
		ID:        t.id,
		Title:     t.title,
		FlightID:  t.flightID,
		Priority:  t.priority,
		Status:    t.status,
		CreatedAt: t.createdAt.Format(time.RFC3339),
	}

	if t.err != nil {
		info.Err = t.err.Error()
	}

	return info
}
