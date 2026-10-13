package main

import (
    "strings"
    "unicode"
)

// ValidatePassword checks if a password meets the security criteria.
func ValidatePassword(password string) bool {
    if len(password) < 8 {
        return false
    }

    if strings.Contains(password, " ") {
        return false
    }

    hasNumber := false
    for _, char := range password {
        if unicode.IsDigit(char) {
            hasNumber = true
            break
        }
    }

    return hasNumber
}
