package main

import (
	"testing"
)

func TestFastPayment_Simple(t *testing.T) {
	g1 := make(chan string, 1)
	g2 := make(chan string)
	g3 := make(chan string)

	expectedMsg := "payment-successful"
	g1 <- expectedMsg

	result, err := FastPayment(g1, g2, g3)

	if err != nil {
		t.Fatalf("Did not expect an error, but got: %v", err)
	}

	if result != expectedMsg {
		t.Errorf("Expected result %q, but got %q", expectedMsg, result)
	}
}
