package main

import (
    "testing"
)

func TestInventory_Sample(t *testing.T) {
    inv := NewInventory()

    inv.AddStock("apple", 10)

    success := inv.Purchase("apple", 4)
    if !success {
        t.Fatalf("Expected purchase to be successful, but it failed")
    }

    remaining := inv.GetStock("apple")
    if remaining != 6 {
        t.Errorf("Expected 6 apples remaining, got %d", remaining)
    }
}
