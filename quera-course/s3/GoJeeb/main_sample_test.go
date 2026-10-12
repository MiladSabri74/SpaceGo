package main

import "testing"

func TestSimpleDeposit(t *testing.T) {
	acc := NewEverydayAccount()

	status := acc.Deposit(100)

	if status != "Success" {
		t.Errorf("Expected status 'Success', got '%s'", status)
	}
	if acc.GetBalance() != 100 {
		t.Errorf("Expected balance to be 100, got %d", acc.GetBalance())
	}
}

func TestSimpleWithdraw(t *testing.T) {
	acc := NewEverydayAccount()
	acc.Deposit(100)

	status := acc.Withdraw(40)

	if status != "Success" {
		t.Errorf("Expected status 'Success', got '%s'", status)
	}
	if acc.GetBalance() != 60 {
		t.Errorf("Expected balance to be 60, got %d", acc.GetBalance())
	}
}
