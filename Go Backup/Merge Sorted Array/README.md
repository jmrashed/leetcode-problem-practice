# Merge Sorted Array Go Project

This is a simple Go project that demonstrates how to merge two sorted arrays into a single sorted array.

## Prerequisites

Make sure you have Go installed on your machine. If not, you can download and install it from the official website: https://golang.org/dl/

## Running the Project

1. Clone or download this repository to your local machine.

2. Open a terminal and navigate to the project directory:

```sh
cd go-leetcode-problem-practice/merge-sorted-array
```

3. Run the Go program using the following command:

```sh
go run main.go
```

4. You should see the merged array output printed in the terminal.

```bash
Merged Array: [1 2 2 3 5 6]
```

# Explain

```go
package main

import (
	"fmt"
)

// merge function takes two slices and two integers as parameters
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1          // Initialize pointer for the last element in nums1
	p2 := n - 1          // Initialize pointer for the last element in nums2
	p := m + n - 1       // Initialize pointer for the last position in nums1

	// Loop while both nums1 and nums2 have elements
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1] // Place the larger element at the end of nums1
			p1--                 // Move p1 and p one step back
		} else {
			nums1[p] = nums2[p2] // Place the larger element at the end of nums1
			p2--                 // Move p2 and p one step back
		}
		p-- // Move p one step back
	}

	// Copy any remaining elements from nums2 to nums1
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

func main() {
	// Initialize the input arrays and sizes
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	// Call the merge function to merge nums2 into nums1
	merge(nums1, m, nums2, n)

	// Print the merged array
	fmt.Printf("Merged Array: %v\n", nums1)
}
```

## Now, let's go through the algorithm step by step:

1. Imports: The program imports the necessary package `fmt` for printing to the console.

2. merge Function: This function takes in two slices (`nums1` and `nums2`) and two integers (m and n) representing the sizes of the slices. It performs the merging operation of the sorted arrays. It follows the same algorithm described earlier:

   - Initialize three pointers: p1 for the last element of `nums1`, p2 for the last element of `nums2`, and p for the last position where elements will be placed in `nums1`.
   - Loop while both `nums1` and `nums2` have elements: Compare the elements at `nums1`[p1] and `nums2`[p2], and place the larger element at `nums1`[p].
   - Move the appropriate pointers (p1, p2, and p) one step back.
   - After the loop, if there are remaining elements in `nums2`, copy them to `nums1` starting from position p.

3. main Function: This is the entry point of the program.

- It initializes the input arrays `nums1` and `nums2` along with their respective sizes m and n.
- Calls the merge function to merge the elements of `nums2` into `nums1`.
- Prints the merged array using `fmt.Printf`.

Overall, the Go code efficiently merges two sorted arrays `nums1` and `nums2` into `nums1`, following the described algorithm, and then outputs the merged array to the console.
