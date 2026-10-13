package main

import (
	"errors"
	"testing"
)

func TestCalculateShipping(t *testing.T) {
	// TODO: Implement table-driven tests for CalculateShipping covering all required scenarios
	usecases := []struct {
		name        string
		weight      float64
		distance    int
		isVIP       bool
		expected    int
		expectedErr error
	}{
		{name: "invalid weight", weight: -1, distance: 10, isVIP: true, expected: 0, expectedErr: errors.New("invalid input")},
		{name: "invalid distance", weight: 10, distance: -2, isVIP: true, expected: 0, expectedErr: errors.New("invalid input")},
		{name: "vip user", weight: 10, distance: 5, isVIP: true, expected: 0, expectedErr: nil},
		{name: "correct value", weight: 10, distance: 10, isVIP: false, expected: 15000, expectedErr: nil},
	}

	for _, tt := range usecases {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := CalculateShipping(tt.weight, tt.distance, tt.isVIP)
			if actual != tt.expected ||
				(actualErr == nil && tt.expectedErr != nil) ||
				(actualErr != nil && tt.expectedErr == nil) ||
				(actualErr != nil && tt.expectedErr != nil &&
					actualErr.Error() != tt.expectedErr.Error()) {
				t.Errorf("CalculateShipping(%f, %d, %v) expected %d, %v but actual is %d, %v",
					tt.weight, tt.distance, tt.isVIP,
					tt.expected, tt.expectedErr,
					actual, actualErr)
			}
		})
	}

}
