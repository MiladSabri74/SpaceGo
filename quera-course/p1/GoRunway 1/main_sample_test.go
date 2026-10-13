package main_test

import (
	"errors"
	"testing"

	"GoRunway/models"
	"GoRunway/task"
)

func TestSample_NewTaskAndExecute(t *testing.T) {
	executed := false
	dummyAction := func() error {
		executed = true
		return nil
	}

	tsk, err := task.NewTask("T-101", "Refueling", "IR-701", models.PriorityHigh, dummyAction)
	if err != nil {
		t.Fatalf("expected no error creating task, got: %v", err)
	}

	if tsk.Status() != models.StatusPending {
		t.Errorf("expected status PENDING, got: %s", tsk.Status())
	}

	err = tsk.Execute()
	if err != nil {
		t.Errorf("expected no error on execution, got: %v", err)
	}

	if !executed {
		t.Errorf("expected action to be executed, but it was not")
	}

	if tsk.Status() != models.StatusCompleted {
		t.Errorf("expected status COMPLETED, got: %s", tsk.Status())
	}
}

func TestSample_ValidationError(t *testing.T) {
	_, err := task.NewTask("", "Refueling", "IR-701", models.PriorityHigh, func() error { return nil })
	if err == nil {
		t.Fatalf("expected error for empty task ID, got nil")
	}

	if !errors.Is(err, models.ErrInvalidTaskID) {
		t.Errorf("expected error to wrap models.ErrInvalidTaskID, got: %v", err)
	}

	var valErr *models.ValidationError
	if !errors.As(err, &valErr) {
		t.Errorf("expected error to be of type *models.ValidationError, got: %T", err)
	}
}
