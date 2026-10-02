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
			line, _ := reader.ReadString('\n')
			fmt.Sscanf(line, "%f %f", &num1, &num2)
			fmt.Printf("Result: %.4f\n", num1+num2)
		} else if instruction == "-" {
			line, _ := reader.ReadString('\n')
			fmt.Sscanf(line, "%f %f", &num1, &num2)
			fmt.Printf("Result: %.4f\n", num1-num2)

		} else if instruction == "*" {
			line, _ := reader.ReadString('\n')
			fmt.Sscanf(line, "%f %f", &num1, &num2)
			fmt.Printf("Result: %.4f\n", num1*num2)

		} else if instruction == "/" {
			line, _ := reader.ReadString('\n')
			fmt.Sscanf(line, "%f %f", &num1, &num2)
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
