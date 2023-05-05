package main

import "fmt"

func twoSum(nums []int, target int) []int {
	t := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, ok := t[v]; ok {
			return []int{vv, i}
		}
		t[target-v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15, 18, 22}
	fmt.Println(twoSum(nums, 29)) // use %v verb to print the slice
}
