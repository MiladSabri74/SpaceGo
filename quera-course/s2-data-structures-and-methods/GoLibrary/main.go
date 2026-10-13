package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Number of commands
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	books := make(map[int]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())

		command := fields[0]
		isbn, _ := strconv.Atoi(fields[1])

		switch command {
		case "add":
			// Everything after ISBN is the title.
			title := strings.Join(fields[2:], " ")
			books[isbn] = title

		case "remove":
			delete(books, isbn)
		}
	}

	// Collect all ISBNs.
	isbns := make([]int, 0, len(books))
	for isbn := range books {
		isbns = append(isbns, isbn)
	}

	// Sort by title first, ISBN second.
	sort.Slice(isbns, func(i, j int) bool {
		isbn1 := isbns[i]
		isbn2 := isbns[j]

		if books[isbn1] != books[isbn2] {
			return books[isbn1] < books[isbn2]
		}

		return isbn1 < isbn2
	})

	// Print sorted ISBNs.
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, isbn := range isbns {
		fmt.Fprintln(writer, isbn)
	}
}
