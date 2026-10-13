package task

import (
	"GoRunway/models"
)

type Task struct {
	// TODO: Define struct fields
}

func NewTask(id, title, flightID string, priority models.Priority, action models.Action) (*Task, error) {
	// TODO: Validate fields and return a new Task in PENDING status or wrapped error
	return nil, nil
}

func (t *Task) ID() string {
	// TODO: Safely return task ID
	return ""
}

func (t *Task) Title() string {
	// TODO: Safely return task title
	return ""
}

func (t *Task) FlightID() string {
	// TODO: Safely return flight ID
	return ""
}

func (t *Task) Priority() models.Priority {
	// TODO: Safely return priority
	return 0
}

func (t *Task) Status() models.TaskStatus {
	// TODO: Safely return status
	return ""
}

func (t *Task) Error() error {
	// TODO: Safely return recorded error
	return nil
}

func (t *Task) Cancel() error {
	// TODO: Cancel pending task or return error
	return nil
}

func (t *Task) Reset() error {
	// TODO: Reset completed/failed/cancelled task to PENDING or return error
	return nil
}

func (t *Task) Execute() error {
	// TODO: Execute action with panic recovery, update status and timestamps
	return nil
}

func (t *Task) GetInfo() models.TaskInfo {
	// TODO: Return immutable snapshot of task info
	return models.TaskInfo{}
}
