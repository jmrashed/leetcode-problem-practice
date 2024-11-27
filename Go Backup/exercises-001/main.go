package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findSevenDivisor(starting int, ending int) string {
	var numbers []string

	for i := starting; i <= ending; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	return strings.Join(numbers, ",")
}
func testFunc() {
	want := "112,119,126,133,147,154,161,168,182,189,196"
	got := findSevenDivisor(100, 200)

	if got != want {
		fmt.Println("Not matching...")
	} else {
		fmt.Println("Matching Successfully")
	}
}
func main() {
	var starting int = 100
	var ending int = 200
	res := findSevenDivisor(starting, ending)
	fmt.Println(res)

	// testFunc()

}
