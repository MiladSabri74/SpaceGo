package main

import (
	"errors"
	"fmt"
)

func Withdraw(balance float64, amount float64) (float64, error) {
	// TODO: Validate amount (< 0 returns error, > balance returns "insufficient funds", else deduct and return new balance)
	if amount <= 0 {
		return 0, fmt.Errorf("invalid amount: %f is negative", amount)
	}
	if balance < amount {
		return 0, errors.New("insufficient funds")
	}
	return balance - amount, nil
}
