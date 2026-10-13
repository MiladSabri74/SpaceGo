package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	votes := make(map[string]int)

	for i := 0; i < n; i++ {
		var name string
		fmt.Fscanln(reader, &name)

		votes[name]++
	}

	winner := ""
	maxVotes := 0

	for name, count := range votes {
		fmt.Printf("%s: %d\n", name, count)

		if count > maxVotes {
			maxVotes = count
			winner = name
		}
	}

	fmt.Println(winner)
}
