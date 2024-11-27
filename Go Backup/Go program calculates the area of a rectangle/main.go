package main

import "fmt"

func main() {
	// Prompt the user to enter the length of the rectangle
	fmt.Println("Enter the length of the rectangle:")

	// Read the length from the user
	var length float64
	fmt.Scanln(&length)

	// Prompt the user to enter the width of the rectangle
	fmt.Println("Enter the width of the rectangle:")

	// Read the width from the user
	var width float64
	fmt.Scanln(&width)

	// Calculate the area of the rectangle
	area := length * width

	// Display the area of the rectangle
	fmt.Printf("The area of the rectangle with length %f and width %f is %f\n", length, width, area)
}
