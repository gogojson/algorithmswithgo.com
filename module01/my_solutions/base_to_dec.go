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
	var val int

	fmt.Printf("VALUE: %d", val)
	var result int
	multiplier := 1

	// charSet := "0123456789ABCDEF"
	for i := len(value) - 1; i >= 0; i-- {
		fmt.Sscanf(string(value[i]), "%X", &val)
		// index := -1
		// for idx, char := range charSet {
		// 	if rune(value[i]) == char {
		// 		index = idx
		// 		break
		// 	}
		// }
		result += val * multiplier
		multiplier *= base
		// fmt.Printf("result: %d, index: %d, multiplyer: %d\n", result, index, multiplier)
	}
	return result
}
