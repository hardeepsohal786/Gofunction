package main

import (
	"fmt"
	"sort"
)

// median calculates the median of a slice of numbers.
// Parameters:
//
//	numbers: A slice of integers.
//
// Returns:
//
//	The median value of the slice. If the slice is empty, returns 0.
func median(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sort.Ints(numbers)
	n := len(numbers)

	if n%2 == 0 {
		// If even, return the average of the two middle numbers.
		mid1 := numbers[n/2-1]
		mid2 := numbers[n/2]
		return float64(mid1+mid2) / 2.0
	}

	// If odd, return the middle number.
	return float64(numbers[n/2])
}

// celsiusToFahrenheit converts a temperature from Celsius to Fahrenheit.
// Parameters:
//
//	celsius: The temperature in Celsius.
//
// Returns:
//
//	The temperature in Fahrenheit.
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	// Test the median function
	nums := []int{3, 1, 4, 1, 5, 9, 2}
	fmt.Printf("Median of %v: %.2f\n", nums, median(nums))

	// Test the celsiusToFahrenheit function
	celsius := 25.0
	fmt.Printf("Temperature in Fahrenheit: %.2f\n", celsiusToFahrenheit(celsius))
}
