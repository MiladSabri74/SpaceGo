package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO: Implement the lab console CLI
	reader := bufio.NewReader(os.Stdin)
	var num1, num2 float64

	for true {
		instruction, _ := reader.ReadString('\n')
		instruction = strings.TrimSpace(instruction)
		if instruction == "+" {
			num1, num2 = readNumbers(reader)
			fmt.Printf("Result: %.4f\n", num1+num2)
		} else if instruction == "-" {
			num1, num2 = readNumbers(reader)
			fmt.Printf("Result: %.4f\n", num1-num2)

		} else if instruction == "*" {
			num1, num2 = readNumbers(reader)
			fmt.Printf("Result: %.4f\n", num1*num2)

		} else if instruction == "/" {
			num1, num2 = readNumbers(reader)
			if num2 != 0 {
				fmt.Printf("Result: %.4f\n", num1/num2)
			} else {
				fmt.Println("Error: Invalid mixture")
			}
		} else if instruction == "exit" {
			fmt.Println("Lab console terminated")
			break
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func readNumbers(reader *bufio.Reader) (float64, float64) {
	line, _ := reader.ReadString('\n')

	var num1, num2 float64
	fmt.Sscanf(line, "%f %f", &num1, &num2)

	return num1, num2
}
