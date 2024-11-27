# Go Program to Subtract Two Integer Numbers
This is a simple Go program that prompts the user to input two integer numbers and then calculates their difference using the subtraction operation. The result is displayed on the screen.

```go 
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
```

In this program, we start by importing the `fmt` package, which provides input/output functionality. Then we define a `main` function which is the entry point of our program.

Inside the `main` function, we use the `fmt.Println` function to prompt the user to enter the first number, and the `fmt.Scanln` function to read the number from the user and store it in the `num1` variable. We repeat the same process for the second number, using the `num2` variable.

Next, we subtract the second number from the first number and store the result in the `result` variable.

Finally, we use the `fmt.Printf` function to display the result, including the original numbers and the subtraction operation.

That's it! This program should allow the user to input two integer numbers and output their difference.