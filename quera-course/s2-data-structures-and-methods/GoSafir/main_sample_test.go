package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSample_GoSafir(t *testing.T) {
	sampleInput := "3\nAli\n40 70 80 20 30\nMaryam\n100 70 30 90 40 70 60\nHana\n10 80 15\n"
	expectedOutput := "Ali Good\nMaryam Very Good\nHana Fair"

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	rIn, wIn, _ := os.Pipe()
	os.Stdin = rIn

	wIn.Write([]byte(sampleInput))
	wIn.Close()

	main()

	wOut.Close()
	var actualOut bytes.Buffer
	io.Copy(&actualOut, rOut)

	expClean := strings.ReplaceAll(strings.TrimSpace(expectedOutput), "\r\n", "\n")
	actClean := strings.ReplaceAll(strings.TrimSpace(actualOut.String()), "\r\n", "\n")

	if expClean != actClean {
		t.Errorf("Sample test mismatch.\nExpected:\n%s\n\nActual:\n%s", expClean, actClean)
	}
}
