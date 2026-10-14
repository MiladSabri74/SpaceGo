package main_test

import (
	"sync"
	"testing"
	"time"

	"GoRunway/models"
	"GoRunway/pool"
	"GoRunway/task"
)

func TestSample_WorkerPoolExecution(t *testing.T) {
	wp, err := pool.NewWorkerPool(3, 10)
	if err != nil {
		t.Fatalf("unexpected error creating worker pool: %v", err)
	}

	if err := wp.Start(); err != nil {
		t.Fatalf("unexpected error starting worker pool: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		act := func() error {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		tsk, _ := task.NewTask("T-SAMPLE", "Sample Action", "FL-100", models.PriorityMedium, act)
		if err := wp.Submit(tsk); err != nil {
			t.Errorf("failed to submit task: %v", err)
		}
	}

	wg.Wait()

	if err := wp.Stop(); err != nil {
		t.Errorf("failed to stop worker pool: %v", err)
	}
}
