package main

import (
    "testing"
)

// Test a standard, successful flow of restocking and selling an item.
func TestVendingMachine_SuccessfulSale(t *testing.T) {
    vm := NewVendingMachine()

    // Setup machine
    vm.TogglePower()      // Turn machine on
    vm.Restock("Soda", 5) // Add 5 sodas

    // Perform a sale
    result := vm.Sell("Soda", 2)

    // Check if the sale succeeded
    if result != "Success" {
        t.Errorf("Expected 'Success', got %q", result)
    }

    // Check if inventory decreased
    if vm.StockCount("Soda") != 4 {
        t.Errorf("Expected 4 sodas left, got %d", vm.StockCount("Soda"))
    }

    // Check if revenue increased
    if vm.TotalRevenue() != 2 {
        t.Errorf("Expected revenue of 2, got %d", vm.TotalRevenue())
    }
}

// Test the two most common reasons a sale would fail.
func TestVendingMachine_BasicFailures(t *testing.T) {
    vm := NewVendingMachine()

    // 1. Try to buy when the machine is OFF
    resultOff := vm.Sell("Chips", 1)
    if resultOff != "Machine is OFF" {
        t.Errorf("Expected 'Machine is OFF', got %q", resultOff)
    }

    // Turn machine ON for the next test
    vm.TogglePower()

    // 2. Try to buy an item that hasn't been restocked
    resultEmpty := vm.Sell("Chips", 1)
    if resultEmpty != "Out of stock" {
        t.Errorf("Expected 'Out of stock', got %q", resultEmpty)
    }
}
