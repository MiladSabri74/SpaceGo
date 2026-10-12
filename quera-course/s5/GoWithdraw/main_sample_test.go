package main

import "testing"

func TestWithdraw_SimpleSuccess(t *testing.T) {
    initialBalance := 100.0
    withdrawalAmount := 30.0
    expectedBalance := 70.0

    newBalance, err := Withdraw(initialBalance, withdrawalAmount)

    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if newBalance != expectedBalance {
        t.Errorf("Expected balance to be %.2f, but got %.2f", expectedBalance, newBalance)
    }
}
