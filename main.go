package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 10
	fmt.Println("Start!")
	var result string
	for i := 0; i < n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			result = result + "Fizz Buzz"
		case i%3 == 0:
			result = result + "Fizz"
		case i%5 == 0:
			result = result + "Buzz"
		default:
			result = result + fmt.Sprint(i)

		}
		if i != n-1 {
			result = result + ", "
		}

	}
	fmt.Println(strings.TrimSpace(result))
}
