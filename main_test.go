package main

import (
	"testing"
)

// TestConvertStringToInt tests the ConvertStringToInt function.
func TestConvertStringToInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
		err      bool
	}{
		{"123", 123, false},
		{"0", 0, false},
		{"-456", -456, false},
		{"abc", 0, true},
		{"", 0, true},
		{"9999999999", 9999999999, false},   // Testing large integer
		{"-9999999999", -9999999999, false}, // Testing large negative integer
	}

	for _, tc := range testCases {
		result, err := ConvertStringToInt(tc.input)
		if (err != nil) != tc.err {
			t.Errorf("ConvertStringToInt(%q) error = %v, expected error = %v", tc.input, err, tc.err)
		}
		if result != tc.expected {
			t.Errorf("ConvertStringToInt(%q) = %d; expected %d", tc.input, result, tc.expected)
		}
	}
}
