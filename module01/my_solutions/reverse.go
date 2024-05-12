package module01

import (
	"fmt"
	"strings"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
// It does the job if the size of the characters is one byte
func ReverseBasic(word string) string {
	result := ""
	for i := len(word) - 1; i >= 0; i-- {
		result = result + string(word[i])
	}
	fmt.Printf("Origin: %s, reversed: %s\n", word, result)
	return result
}

// Check out the strings.Builder library it is a cool way to create strings more efficiently
func ReverseWithEfficiency(word string) string {
	sb := strings.Builder{}
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	fmt.Printf("Origin: %s, reversed: %s\n", word, sb.String())
	return sb.String()
}

func Reverse(word string) string {
	result := ""
	for _, v := range word {
		result = string(v) + result
	}
	fmt.Printf("Origin: %s, reversed: %s\n", word, result)
	return result
}
