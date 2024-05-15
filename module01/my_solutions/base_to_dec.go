package module01

import "fmt"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	charSet := "0123456789ABCDEF"
	var result int
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for idx, char := range charSet {
			if rune(value[i]) == char {
				index = idx
				break
			}
		}
		result += index * multiplier
		multiplier *= base
		fmt.Printf("result: %d, index: %d, multiplyer: %d\n", result, index, multiplier)
	}
	return result
}
