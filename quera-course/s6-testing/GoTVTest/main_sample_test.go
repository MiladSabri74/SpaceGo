package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const sampleReferenceImplementation = `package main
import (
	"errors"
	"sync"
)
type Poll struct {
	votes map[string]int
	mu    sync.Mutex
}
func NewPoll() *Poll {
	return &Poll{
		votes: make(map[string]int),
	}
}
func (p *Poll) CastVote(option string) error {
	if option == "" {
		return errors.New("empty option")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.votes[option]++
	return nil
}
func (p *Poll) GetResults() map[string]int {
	p.mu.Lock()
	defer p.mu.Unlock()
	results := make(map[string]int, len(p.votes))
	for k, v := range p.votes {
		results[k] = v
	}
	return results
}
`

const sampleBuggyNoMutex = `package main
import (
	"errors"
	"time"
)
type Poll struct {
	votes map[string]int
}
func NewPoll() *Poll {
	return &Poll{
		votes: make(map[string]int),
	}
}
func (p *Poll) CastVote(option string) error {
	if option == "" {
		return errors.New("empty option")
	}
	current := p.votes[option]
	time.Sleep(1 * time.Millisecond)
	p.votes[option] = current + 1
	return nil
}
func (p *Poll) GetResults() map[string]int {
	results := make(map[string]int, len(p.votes))
	for k, v := range p.votes {
		results[k] = v
	}
	return results
}
`

func runSampleGoTest(t *testing.T, files map[string]string) ([]byte, error) {
	t.Helper()
	tempDir := t.TempDir()

	goModBytes, err := os.ReadFile("go.mod")
	if err == nil {
		os.WriteFile(filepath.Join(tempDir, "go.mod"), goModBytes, 0644)
	}
	goSumBytes, err := os.ReadFile("go.sum")
	if err == nil {
		os.WriteFile(filepath.Join(tempDir, "go.sum"), goSumBytes, 0644)
	}

	for name, content := range files {
		err := os.WriteFile(filepath.Join(tempDir, name), []byte(content), 0644)
		if err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	cmd := exec.Command("go", "test", "-count=1", ".")
	cmd.Dir = tempDir
	return cmd.CombinedOutput()
}

func TestSample_StudentTestsPassReference(t *testing.T) {
	studentTestCodeBytes, err := os.ReadFile("poll_test.go")
	if err != nil {
		t.Fatalf("read submitted test file: %v", err)
	}
	studentTestCode := string(studentTestCodeBytes)

	out, err := runSampleGoTest(t, map[string]string{
		"poll.go":      sampleReferenceImplementation,
		"poll_test.go": studentTestCode,
	})
	if err != nil {
		t.Errorf("Your tests (poll_test.go) failed when running against a CORRECT implementation!\nOutput:\n%s", out)
	}
}

func TestSample_BugDetection(t *testing.T) {
	studentTestCodeBytes, err := os.ReadFile("poll_test.go")
	if err != nil {
		t.Fatalf("read submitted test file: %v", err)
	}
	studentTestCode := string(studentTestCodeBytes)

	out, err := runSampleGoTest(t, map[string]string{
		"poll.go":      sampleBuggyNoMutex,
		"poll_test.go": studentTestCode,
	})
	if err == nil {
		t.Errorf("Your tests PASSED against a BUGGY implementation. You missed the concurrent race condition bug!")
	} else if !strings.Contains(string(out), "concurrent map writes") && !strings.Contains(string(out), "FAIL") {
		// They must trigger the race condition which will cause a panic or a failure
		t.Errorf("Your tests failed, but not for the right reason. Expected race condition.")
	}
}
