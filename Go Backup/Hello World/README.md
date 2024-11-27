# "Hello World" program in Go
This is a simple "Hello World" program written in Go. It demonstrates how to print a message to the console using the `fmt` package.
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}

```
In this program, we first declare the `package main`, which indicates that this is the main package of our program. We then import the `fmt` package, which provides functions for formatting and printing strings.

Inside the `main()` function, we use the `fmt.Println()` function to print the string "Hello, world!" to the console. When we run this program, we should see the output "Hello, world!" printed on the screen.

Note that in Go, every statement must end with a semicolon, although it's common to omit them when the statement is followed by a newline. Also, Go is a statically typed language, which means that we need to declare the type of each variable we use. However, since we're not using any variables in this program, we don't need to declare any types.

## Requirements
To run this program, you need to have Go installed on your computer. You can download Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

## Usage
1. Clone this repository to your local machine.
2. Open a terminal and navigate to the directory where the repository is located.
3. Run the command `go run '.\Hello World.go'` to compile and run the program.
4. The program will print the message "Hello, world!" to the console.

## Contributing
If you find a bug or want to suggest an improvement, please open an issue or submit a pull request.
