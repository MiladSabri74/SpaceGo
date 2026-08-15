package manager

import (
	"fmt"
	"os"

	"github.com/spacego/file-manager/internal/fileops"
)

// Manager wraps the fileops.Manager and provides CLI-friendly methods.
type Manager struct {
	ops fileops.FileOperator
}

// New creates a new Manager instance.
func New() *Manager {
	return &Manager{
		ops: fileops.NewManager(),
	}
}

// ListFiles lists files in the given directory and prints them to stdout.
func (m *Manager) ListFiles(dir string) error {
	names, err := m.ops.List(dir)
	if err != nil {
		return err
	}

	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}

// CopyFile copies a file from src to dst.
func (m *Manager) CopyFile(src, dst string) error {
	return m.ops.Copy(src, dst)
}

// FileExists checks if a file or directory exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDirectory checks if the given path is a directory.
func IsDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
