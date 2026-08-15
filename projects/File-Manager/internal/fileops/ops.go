package fileops

import (
	"fmt"
	"io"
	"os"
)

// FileOperator defines the interface for file operations.
type FileOperator interface {
	List(dir string) ([]string, error)
	Copy(src, dst string) error
}

// Manager implements FileOperator interface.
type Manager struct{}

// NewManager creates a new Manager instance.
func NewManager() *Manager {
	return &Manager{}
}

// List returns a list of file and directory names in the given directory.
func (m *Manager) List(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %q: %w", dir, err)
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names, nil
}

// Copy copies the source file to the destination file.
func (m *Manager) Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("opening source file %q: %w", src, err)
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return fmt.Errorf("stating source file %q: %w", src, err)
	}

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return fmt.Errorf("creating destination file %q: %w", dst, err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copying data from %q to %q: %w", src, dst, err)
	}

	if err := out.Close(); err != nil {
		return fmt.Errorf("closing destination file %q: %w", dst, err)
	}

	return nil
}
