package main

import (
	"errors"
	"fmt"
)

var ErrDBConnection = errors.New("database connection lost")

func insertRecord(username string) error {
	if username == "hacker" {
		return ErrDBConnection
	}
	return nil
}

func ValidateUser(age int, username string) error {
	// TODO: Validate age (>= 18) and non-empty username, join errors with errors.Join, or return nil
	var errs []error
	if age < 18 {
		errs = append(errs, errors.New("age must be at least 18"))
	}
	if username == "" {
		errs = append(errs, errors.New("username cannot be empty"))
	}
	return errors.Join(errs...)
}

func SaveUser(username string) error {
	// TODO: Call insertRecord and wrap any error using fmt.Errorf with "failed to save user: %w"
	if err := insertRecord(username); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func CheckDBError(err error) bool {
	// TODO: Check if ErrDBConnection is present in err chain using errors.Is
	return errors.Is(err, ErrDBConnection)
}
