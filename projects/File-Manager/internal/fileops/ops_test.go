package fileops

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewManager(t *testing.T) {
	mgr := NewManager()
	if mgr == nil {
		t.Fatal("NewManager() returned nil")
	}
}

func TestList(t *testing.T) {
	mgr := NewManager()

	// Create a temporary directory with some files
	tmpDir, err := os.MkdirTemp("", "fileops_test_list_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	testFiles := []string{"file1.txt", "file2.txt", "subdir"}
	for _, name := range testFiles {
		path := filepath.Join(tmpDir, name)
		if name == "subdir" {
			if err := os.Mkdir(path, 0755); err != nil {
				t.Fatalf("Failed to create subdir: %v", err)
			}
		} else {
			if err := os.WriteFile(path, []byte("test"), 0644); err != nil {
				t.Fatalf("Failed to create file %s: %v", name, err)
			}
		}
	}

	// Test listing directory
	names, err := mgr.List(tmpDir)
	if err != nil {
		t.Fatalf("List() returned error: %v", err)
	}

	if len(names) != len(testFiles) {
		t.Errorf("Expected %d files, got %d", len(testFiles), len(names))
	}

	// Verify all expected files are present
	fileSet := make(map[string]bool)
	for _, name := range names {
		fileSet[name] = true
	}

	for _, expected := range testFiles {
		if !fileSet[expected] {
			t.Errorf("Expected file %q not found in list", expected)
		}
	}
}

func TestListNonExistentDirectory(t *testing.T) {
	mgr := NewManager()

	_, err := mgr.List("/nonexistent/directory/path")
	if err == nil {
		t.Error("Expected error for non-existent directory, got nil")
	}
}

func TestCopy(t *testing.T) {
	mgr := NewManager()

	// Create temporary directory
	tmpDir, err := os.MkdirTemp("", "fileops_test_copy_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create source file with known content
	srcPath := filepath.Join(tmpDir, "source.txt")
	content := []byte("Hello, World! This is test content.")
	if err := os.WriteFile(srcPath, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Define destination path
	dstPath := filepath.Join(tmpDir, "destination.txt")

	// Test copy operation
	err = mgr.Copy(srcPath, dstPath)
	if err != nil {
		t.Fatalf("Copy() returned error: %v", err)
	}

	// Verify destination file exists
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		t.Fatal("Destination file was not created")
	}

	// Verify content matches
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstContent) != string(content) {
		t.Errorf("Content mismatch. Expected %q, got %q", string(content), string(dstContent))
	}

	// Verify file permissions are preserved
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		t.Fatalf("Failed to stat source file: %v", err)
	}

	dstInfo, err := os.Stat(dstPath)
	if err != nil {
		t.Fatalf("Failed to stat destination file: %v", err)
	}

	if srcInfo.Mode() != dstInfo.Mode() {
		t.Errorf("Permission mismatch. Source: %o, Destination: %o", srcInfo.Mode(), dstInfo.Mode())
	}
}

func TestCopyNonExistentSource(t *testing.T) {
	mgr := NewManager()

	tmpDir, err := os.MkdirTemp("", "fileops_test_copy_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	srcPath := filepath.Join(tmpDir, "nonexistent.txt")
	dstPath := filepath.Join(tmpDir, "destination.txt")

	err = mgr.Copy(srcPath, dstPath)
	if err == nil {
		t.Error("Expected error for non-existent source, got nil")
	}
}

func TestCopyToInvalidDestination(t *testing.T) {
	mgr := NewManager()

	tmpDir, err := os.MkdirTemp("", "fileops_test_copy_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create source file
	srcPath := filepath.Join(tmpDir, "source.txt")
	if err := os.WriteFile(srcPath, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Try to copy to invalid path
	dstPath := "/nonexistent/directory/destination.txt"

	err = mgr.Copy(srcPath, dstPath)
	if err == nil {
		t.Error("Expected error for invalid destination, got nil")
	}
}

func TestCopyOverwritesExistingFile(t *testing.T) {
	mgr := NewManager()

	tmpDir, err := os.MkdirTemp("", "fileops_test_copy_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create source file
	srcPath := filepath.Join(tmpDir, "source.txt")
	srcContent := []byte("New content")
	if err := os.WriteFile(srcPath, srcContent, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Create destination file with different content
	dstPath := filepath.Join(tmpDir, "destination.txt")
	oldContent := []byte("Old content")
	if err := os.WriteFile(dstPath, oldContent, 0644); err != nil {
		t.Fatalf("Failed to create destination file: %v", err)
	}

	// Copy should overwrite
	err = mgr.Copy(srcPath, dstPath)
	if err != nil {
		t.Fatalf("Copy() returned error: %v", err)
	}

	// Verify content was overwritten
	newContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(newContent) != string(srcContent) {
		t.Errorf("File was not properly overwritten. Expected %q, got %q", string(srcContent), string(newContent))
	}
}
