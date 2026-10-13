package main

import "testing"

func TestLogin_Success(t *testing.T) {
	mockDB := map[string]string{
		"alice": "active",
	}
	username := "alice"

	err := Login(mockDB, username)

	if err != nil {
		t.Errorf("Expected no error for active user %q, but got: %v", username, err)
	}
}
