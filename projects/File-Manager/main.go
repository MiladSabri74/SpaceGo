package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		show_help(1)
	}

	command := strings.ToLower(os.Args[1])
	switch command {
	case "list":
		list(os.Args[2])
	case "copy":
		if len(os.Args) != 4 {
			show_help(1)
		}
		if err := copy(os.Args[2], os.Args[3]); err != nil {
			fmt.Fprintf(os.Stderr, "copy: %v\n", err)
			os.Exit(1)
		}
	default:
		show_help(0)
	}
}

func list(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		handle_error("list", err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy: %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		handle_error("copy", err)
	}

	out, err := os.OpenFile(
		dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		handle_error("copy", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		handle_error("copy", err)
	}
	return out.Close()
}

func show_help(exit_status int) {
	help_message := "File-Manager is a CLI tool for managing files:\n\n" +
		"Usage:\n" +
		"List <path>\tShow list of files in path\n" +
		"help\t\tShow list of supported command and usages"

	fmt.Println(help_message)
	os.Exit(exit_status)
}

func handle_error(cmd string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", cmd, err)
	os.Exit(1)
}
