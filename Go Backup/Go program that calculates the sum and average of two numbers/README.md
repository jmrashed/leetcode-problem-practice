# Go Program to Calculate the Sum and Average of Two Numbers
This is a simple Go program that prompts the user to input two numbers, calculates their sum and average, and then displays the results on the screen.

```go
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

```


In this program, we start by importing the `fmt` package for input/output functionality.

Inside the main function, we use the `fmt.Println` function to prompt the user to enter the first number, and the `fmt.Scanln` function to read the number from the user and store it in the `num1` variable. We repeat the same process for the second number, using the `num2` variable.

Next, we calculate the sum of the two numbers and store it in the `sum` variable. We calculate the average by dividing the sum by `2.0` and store it in the average variable.

Finally, we use the `fmt.Printf` function to display the sum and average of the two numbers, including the original numbers in the output.

That's it! This program should allow the user to input two numbers and output their `sum` and `average`.