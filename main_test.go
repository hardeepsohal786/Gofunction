package main

import (
	"testing"
)

// TestMedian tests the median function for various inputs.
func TestMedian(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{7, 5, 3, 1, 9}, 5.0},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := median(test.numbers)
		if result != test.expected {
			t.Errorf("median(%v) = %.2f; expected %.2f", test.numbers, result, test.expected)
		}
	}
}

// TestCelsiusToFahrenheit tests the celsiusToFahrenheit function for various inputs.
func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius  float64
		expected float64
	}{
		{0, 32},
		{25, 77},
		{-10, 14},
		{100, 212},
	}

	for _, test := range tests {
		result := celsiusToFahrenheit(test.celsius)
		if result != test.expected {
			t.Errorf("celsiusToFahrenheit(%.2f) = %.2f; expected %.2f", test.celsius, result, test.expected)
		}
	}
}
