package main

import "fmt"

func main() {
	// Prompt the user to enter the first number
	fmt.Println("Enter the first number:")

	// Read the first number from the user
	var num1 float64
	fmt.Scanln(&num1)

	// Prompt the user to enter the second number
	fmt.Println("Enter the second number:")

	// Read the second number from the user
	var num2 float64
	fmt.Scanln(&num2)

	// Calculate the sum of the two numbers
	sum := num1 + num2

	// Calculate the average of the two numbers
	average := sum / 2.0

	// Display the sum and average of the two numbers
	fmt.Printf("The sum of %f and %f is %f\n", num1, num2, sum)
	fmt.Printf("The average of %f and %f is %f\n", num1, num2, average)
}
