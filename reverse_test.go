
package main

import (
    "testing"
)

func TestReverse(t *testing.T) {
    tests := []struct {
        input, expected string
    }{
        {"hello", "olleh"},
        {"world", "dlrow"},
        {"", ""},
        {"a", "a"},
        {"Go", "oG"},
    }

    for _, test := range tests {
        result := Reverse(test.input)
        if result != test.expected {
            t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
        }
    }
}
