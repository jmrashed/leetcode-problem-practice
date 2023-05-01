package main

import "fmt"

func main() {
	// Prompt user to enter the first number
	fmt.Println("Enter the first number: ")

	// Read the first number from the user
	var num1 int
	fmt.Scanln(&num1)

	// Prompt user to enter the second number
	fmt.Println("Enter the second number: ")

	// Read the second number from the user
	var num2 int
	fmt.Scanln(&num2)

	// Subtract the second number from the first number
	result := num1 - num2

	// Display the result
	fmt.Printf("The subtraction of %d and %d is %d\n", num1, num2, result)
}
