package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spacego/file-manager/pkg/manager"
)

func main() {
	if len(os.Args) < 2 {
		showHelp(1)
	}

	command := strings.ToLower(os.Args[1])
	mgr := manager.New()

	switch command {
	case "list":
		if len(os.Args) != 3 {
			fmt.Fprintln(os.Stderr, "Error: 'list' command requires exactly one argument (directory path)")
			showHelp(1)
		}
		dir := os.Args[2]
		if err := mgr.ListFiles(dir); err != nil {
			handleError("list", err)
		}

	case "copy":
		if len(os.Args) != 4 {
			fmt.Fprintln(os.Stderr, "Error: 'copy' command requires exactly two arguments (source and destination)")
			showHelp(1)
		}
		src := os.Args[2]
		dst := os.Args[3]
		if err := mgr.CopyFile(src, dst); err != nil {
			handleError("copy", err)
		}

	case "help":
		showHelp(0)

	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command %q\n", command)
		showHelp(1)
	}
}

func showHelp(exitStatus int) {
	helpMessage := `File-Manager is a CLI tool for managing files.

Usage:
  file-manager <command> [arguments]

Commands:
  list <path>           Show list of files in the specified path
  copy <src> <dst>      Copy source file to destination file
  help                  Show this help message

Examples:
  file-manager list /home/user/documents
  file-manager copy source.txt destination.txt
`
	fmt.Println(helpMessage)
	os.Exit(exitStatus)
}

func handleError(cmd string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
	os.Exit(1)
}
