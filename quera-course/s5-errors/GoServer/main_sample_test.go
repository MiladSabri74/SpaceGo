package main

import (
	"testing"
)

func TestSafeExecutor_BasicPanic(t *testing.T) {
	// Arrange
	expectedErrMsg := "something went terribly wrong"
	panickingTask := func() {
		panic(expectedErrMsg)
	}

	// Act
	err := SafeExecutor(panickingTask)

	// Assert
	if err == nil {
		t.Fatal("Expected an error to be returned, but got nil")
	}

	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message %q, but got %q", expectedErrMsg, err.Error())
	}
}
