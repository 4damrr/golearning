package main

import (
	"fmt"
	"os"
)

func getUserInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return input
}

func main() {
	num1 := getUserInput("Enter your first number: ")
	num2 := getUserInput("Enter your second number: ")
	fmt.Print(num1)
	fmt.Print(num2)
}
