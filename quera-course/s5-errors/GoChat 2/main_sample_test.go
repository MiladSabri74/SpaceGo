package main

import (
	"testing"
)

func TestInsertRecordSample(t *testing.T) {
	err := insertRecord("alice")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestValidateUserSample(t *testing.T) {
	err := ValidateUser(25, "john_doe")
	if err != nil {
		t.Errorf("Expected valid user to pass, got error: %v", err)
	}
}

func TestSaveUserSample(t *testing.T) {
	err := SaveUser("bob")
	if err != nil {
		t.Errorf("Expected user to save successfully, got error: %v", err)
	}
}

func TestCheckDBErrorSample(t *testing.T) {
	isDBErr := CheckDBError(ErrDBConnection)
	if !isDBErr {
		t.Errorf("Expected true for ErrDBConnection, got false")
	}
}
