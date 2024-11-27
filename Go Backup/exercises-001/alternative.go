package main

import "fmt"

func main() {
	var nums []int
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			nums = append(nums, i)
		}
	}
	fmt.Printf("%v", nums[0])
	for i := 1; i < len(nums); i++ {
		fmt.Printf(",%v", nums[i])
	}
	fmt.Println()
}
