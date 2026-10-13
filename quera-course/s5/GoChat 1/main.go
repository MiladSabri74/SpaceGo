package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found") // TODO: Define sentinel error with message "user not found" using errors.New

type AuthError struct {
	// TODO: Define Username and Reason string fields
	username string
	reason   string
}

func (e AuthError) Error() string {
	// TODO: Format and return "authentication failed for user <Username>: <Reason>"
	return fmt.Sprintf("authentication failed for user %s: %s", e.username, e.reason)
}

func Login(db map[string]string, username string) error {
	// TODO: Check user status in db; return ErrUserNotFound if missing, AuthError if blocked, or nil if active
	if status, exists := db[username]; exists {
		if status == "active" {
			return nil
		}
		return AuthError{username: username, reason: "account is blocked"}
	}
	return ErrUserNotFound
}
