package main

import "testing"

// TestIsEven tests the IsEven function with various test cases.
func TestIsEven(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
		{-3, false},
	}

	for _, tc := range testCases {
		result := IsEven(tc.n)
		if result != tc.expected {
			t.Errorf("IsEven(%d) = %v; expected %v", tc.n, result, tc.expected)
		}
	}
}
