
package main

// Reverse returns the reverse of the given string.
func Reverse(s string) string {
    reversed := ""
    for _, char := range s {
        reversed = string(char) + reversed
    }
    return reversed
}
