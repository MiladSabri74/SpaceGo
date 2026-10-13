package main

import (
	"testing"
)

func TestValidatePassword(t *testing.T) {
	// TODO: Implement test cases for ValidatePassword function
	actual := ValidatePassword("MyPassword@1374")
	expected := true
	if actual != expected {
		t.Errorf("expected true but return false")
	}

	actual = ValidatePassword("myNewPassword")
	expected = false
	if actual != expected {
		t.Errorf("expected false. without any numeric character")
	}

	actual = ValidatePassword("my newpassword1")
	expected = false
	if actual != expected {
		t.Errorf("expected false. invalid spcace character")
	}

	actual = ValidatePassword("my123")
	expected = false
	if actual != expected {
		t.Errorf("expected false. short password")
	}

}
