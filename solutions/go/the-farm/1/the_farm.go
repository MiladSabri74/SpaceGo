package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	amount, err := calc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount / float64(cows) * factor, nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calc, cows)
}

type InvalidCowsError struct {
	Cows    int
	Message string
}

func (err InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.Cows, err.Message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{Cows: cows, Message: "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{Cows: cows, Message: "no cows don't need food"}
	}
	return nil
}
