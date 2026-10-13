package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSample_GoLibrary(t *testing.T) {
	sampleInput := "8\nadd 9789643110543 The Catcher in the Rye\nadd 9786001191612 Animal Farm\nadd 9789641740926 The Great Gatsby\nadd 9780747532743 Harry Potter and the Philosophers Stone\nadd 9786001192084 Animal Farm\nremove 9789641740926\nremove 1234567890\nadd 9780061120084 To Kill a Mockingbird\n"
	expectedOutput := "9786001191612\n9786001192084\n9780747532743\n9789643110543\n9780061120084"

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
