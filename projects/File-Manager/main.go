package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type FileOperator interface {
	list(dir string) error
	copy(src, dst string) error
}

type Manager struct{}

func main() {

	if len(os.Args) < 3 {
		showHelp(1)
	}

	command := strings.ToLower(os.Args[1])
	manager := &Manager{}

	switch command {
	case "list":
		if err := manager.list(os.Args[2]); err != nil {
			handleError(command, err)
		}
	case "copy":
		if len(os.Args) != 4 {
			showHelp(1)
		}
		if err := manager.copy(os.Args[2], os.Args[3]); err != nil {
			handleError(command, err)
		}
	default:
		showHelp(0)
	}
}

func (m *Manager) list(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	return nil
}

func (m *Manager) copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(
		dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

func showHelp(exit_status int) {
	help_message := "File-Manager is a CLI tool for managing files:\n\n" +
		"Usage:\n" +
		"list <path>\t\tShow list of files in path\n" +
		"copy <src> <dst>\tCopy src file to dst file\n" +
		"help\t\t\tShow list of supported command and usages"

	fmt.Println(help_message)
	os.Exit(exit_status)
}

func handleError(cmd string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
	os.Exit(1)
}
