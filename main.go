package main

import (
	"fmt"
	"strconv"
)

// ConvertStringToInt converts a string to an integer. Returns an error if the conversion fails.
func ConvertStringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func main() {
	// Example usage
	strs := []string{"123", "456", "789", "invalid"}
	for _, str := range strs {
		val, err := ConvertStringToInt(str)
		if err != nil {
			fmt.Printf("Error converting '%s': %v\n", str, err)
		} else {
			fmt.Printf("Converted '%s' to %d\n", str, val)
		}
	}
}
