package main

import (
	"fmt"
)

// IsEven returns true if the number is even, false otherwise.
func IsEven(n int) bool {
	return n%2 == 0
}

func main() {
	// Example usage of the IsEven function
	numbers := []int{2, 3, 0, -2, -3}
	for _, n := range numbers {
		fmt.Printf("IsEven(%d) = %v\n", n, IsEven(n))
	}
}
