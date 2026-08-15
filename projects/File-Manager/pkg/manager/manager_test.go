package manager

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	mgr := New()
	if mgr == nil {
		t.Fatal("New() returned nil")
	}
}

func TestFileExists(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "manager_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpPath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpPath)

	// Test existing file
	if !FileExists(tmpPath) {
		t.Error("FileExists() returned false for existing file")
	}

	// Test non-existing file
	nonExistentPath := filepath.Join(filepath.Dir(tmpPath), "nonexistent_file.txt")
	if FileExists(nonExistentPath) {
		t.Error("FileExists() returned true for non-existent file")
	}
}

func TestIsDirectory(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "manager_test_dir_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "manager_test_file_*")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpFilePath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpFilePath)

	// Test directory
	isDir, err := IsDirectory(tmpDir)
	if err != nil {
		t.Fatalf("IsDirectory() returned error for valid directory: %v", err)
	}
	if !isDir {
		t.Error("IsDirectory() returned false for a directory")
	}

	// Test file
	isDir, err = IsDirectory(tmpFilePath)
	if err != nil {
		t.Fatalf("IsDirectory() returned error for valid file: %v", err)
	}
	if isDir {
		t.Error("IsDirectory() returned true for a file")
	}

	// Test non-existent path
	_, err = IsDirectory("/nonexistent/path")
	if err == nil {
		t.Error("IsDirectory() should return error for non-existent path")
	}
}

func TestListFiles(t *testing.T) {
	mgr := New()

	// Create a temporary directory with test files
	tmpDir, err := os.MkdirTemp("", "manager_test_list_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	testFiles := []string{"test1.txt", "test2.txt", "test3.txt"}
	for _, name := range testFiles {
		path := filepath.Join(tmpDir, name)
		if err := os.WriteFile(path, []byte("test content"), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", name, err)
		}
	}

	// Test ListFiles - this will print to stdout, so we mainly test it doesn't error
	err = mgr.ListFiles(tmpDir)
	if err != nil {
		t.Errorf("ListFiles() returned error: %v", err)
	}

	// Test with non-existent directory
	err = mgr.ListFiles("/nonexistent/directory")
	if err == nil {
		t.Error("ListFiles() should return error for non-existent directory")
	}
}

func TestCopyFile(t *testing.T) {
	mgr := New()

	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "manager_test_copy_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create source file
	srcPath := filepath.Join(tmpDir, "source.txt")
	content := []byte("Test content for copy operation")
	if err := os.WriteFile(srcPath, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Define destination
	dstPath := filepath.Join(tmpDir, "destination.txt")

	// Test CopyFile
	err = mgr.CopyFile(srcPath, dstPath)
	if err != nil {
		t.Fatalf("CopyFile() returned error: %v", err)
	}

	// Verify file was copied
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstContent) != string(content) {
		t.Errorf("CopyFile() did not preserve content. Expected %q, got %q", string(content), string(dstContent))
	}

	// Test with non-existent source
	err = mgr.CopyFile("/nonexistent/source.txt", dstPath)
	if err == nil {
		t.Error("CopyFile() should return error for non-existent source")
	}
}
