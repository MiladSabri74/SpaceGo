package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const sampleReferenceImplementation = `package main

import (
	"strings"
	"unicode"
)

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if strings.Contains(password, " ") {
		return false
	}
	for _, char := range password {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
`

const sampleMissingLengthCheck = `package main

import (
	"strings"
	"unicode"
)

func ValidatePassword(password string) bool {
	if strings.Contains(password, " ") {
		return false
	}
	for _, char := range password {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
`

func runSampleStudentTests(t *testing.T, implementation string) ([]byte, error) {
	t.Helper()

	studentTestCode, err := os.ReadFile("auth_test.go")
	if err != nil {
		t.Fatalf("read submitted test file: %v", err)
	}

	testDirectory := t.TempDir()
	files := map[string][]byte{
		"go.mod":       []byte("module evaluator\n\ngo 1.26\n"),
		"auth.go":      []byte(implementation),
		"auth_test.go": studentTestCode,
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(testDirectory, name), content, 0o600); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}

	command := exec.Command("go", "test", "-count=1", ".")
	command.Dir = testDirectory
	return command.CombinedOutput()
}

func TestSample_StudentTestsPassReference(t *testing.T) {
	if output, err := runSampleStudentTests(t, sampleReferenceImplementation); err != nil {
		t.Fatalf("submitted tests must pass against the correct implementation:\n%s", output)
	}
}

func TestSample_StudentTestsCatchMissingLengthCheck(t *testing.T) {
	if output, err := runSampleStudentTests(t, sampleReferenceImplementation); err != nil {
		t.Fatalf("submitted tests must pass against the correct implementation:\n%s", output)
	}
	if output, err := runSampleStudentTests(t, sampleMissingLengthCheck); err == nil {
		t.Fatalf("Your tests did not reject a password shorter than 8 characters.\nSubmitted test output:\n%s", output)
	}
}
