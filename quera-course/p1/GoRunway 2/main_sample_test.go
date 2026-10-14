package main_test

import (
	"errors"
	"testing"

	"GoRunway/models"
	"GoRunway/repository"
	"GoRunway/task"
)

func TestSample_RepositorySaveAndFind(t *testing.T) {
	repo := repository.NewTaskRepository()

	dummyAct := func() error { return nil }
	tsk, err := task.NewTask("T-1", "Baggage Loading", "IR-101", models.PriorityHigh, dummyAct)
	if err != nil {
		t.Fatalf("unexpected error creating task: %v", err)
	}

	err = repo.Save(tsk)
	if err != nil {
		t.Fatalf("expected no error saving task, got: %v", err)
	}

	found, err := repo.FindByID("T-1")
	if err != nil {
		t.Fatalf("expected to find task T-1, got error: %v", err)
	}

	if found.Title() != "Baggage Loading" {
		t.Errorf("expected Title 'Baggage Loading', got '%s'", found.Title())
	}

	_, err = repo.FindByID("NON-EXISTENT")
	if err == nil || !errors.Is(err, repository.ErrTaskNotFound) {
		t.Errorf("expected ErrTaskNotFound for non-existent ID, got: %v", err)
	}
}
