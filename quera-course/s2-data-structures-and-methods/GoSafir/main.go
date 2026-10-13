package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRank(avg float64) string {
	if avg >= 80 {
		return "Excellent"
	} else if avg >= 60 {
		return "Very Good"
	} else if avg >= 40 {
		return "Good"
	}

	return "Fair"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		line, _ := reader.ReadString('\n')
		scores := strings.Fields(line)

		sum := 0

		for _, score := range scores {
			value, _ := strconv.Atoi(score)
			sum += value
		}

		avg := float64(sum) / float64(len(scores))

		fmt.Fprintln(writer, name, getRank(avg))
	}
}
