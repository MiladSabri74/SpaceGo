package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const sampleReferenceImplementation = `package main
import "errors"
func CalculateShipping(weight float64, distance int, isVIP bool) (int, error) {
	if weight < 0 || distance < 0 {
		return 0, errors.New("invalid input")
	}
	if isVIP {
		return 0, nil
	}
	return int(weight*1000) + (distance * 500), nil
}
`

const sampleBuggyNoVIP = `package main
import "errors"
func CalculateShipping(weight float64, distance int, isVIP bool) (int, error) {
	if weight < 0 || distance < 0 {
		return 0, errors.New("invalid input")
	}
	return int(weight*1000) + (distance * 500), nil
}
`

func runSampleStudentTests(t *testing.T, implementation string) ([]byte, error) {
	t.Helper()

	studentTestCode, err := os.ReadFile("shipping_test.go")
	if err != nil {
		t.Fatalf("read submitted test file: %v", err)
	}

	tempDir := t.TempDir()

	err = os.WriteFile(filepath.Join(tempDir, "shipping_test.go"), studentTestCode, 0644)
	if err != nil {
		t.Fatalf("write student tests: %v", err)
	}

	err = os.WriteFile(filepath.Join(tempDir, "shipping.go"), []byte(implementation), 0644)
	if err != nil {
		t.Fatalf("write implementation: %v", err)
	}

	err = os.WriteFile(filepath.Join(tempDir, "go.mod"), []byte("module grading\n\ngo 1.26\n"), 0644)
	if err != nil {
		t.Fatalf("write go.mod: %v", err)
	}

	cmd := exec.Command("go", "test", "-count=1", ".")
	cmd.Dir = tempDir
	return cmd.CombinedOutput()
}

func requireSampleStudentTestsPass(t *testing.T, output []byte, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Your tests failed when running against a CORRECT implementation!\nOutput:\n%s", output)
	}
}

func requireSampleStudentTestsFail(t *testing.T, output []byte, err error, bugMessage string) {
	t.Helper()
	if err == nil {
		t.Errorf("Your tests PASSED against a BUGGY implementation. You missed a bug! %s", bugMessage)
	}
}

func TestSample_StudentTestsPassReference(t *testing.T) {
	out, err := runSampleStudentTests(t, sampleReferenceImplementation)
	requireSampleStudentTestsPass(t, out, err)
}

func TestSample_StudentTestsCatchNoVIP(t *testing.T) {
	out, err := runSampleStudentTests(t, sampleBuggyNoVIP)
	requireSampleStudentTestsFail(t, out, err, "The code ignored VIP status but your tests didn't catch it.")
}
