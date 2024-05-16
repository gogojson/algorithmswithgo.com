package module01

import (
	"fmt"
	"strconv"
)

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//	BaseToBase("E", 16, 2) => "1110"
func BaseToBase(value string, base, newBase int) string {
	switch {
	case newBase == 10:
		return fmt.Sprint(BaseToDec(value, base))
	case base == 10:
		i, err := strconv.Atoi(value)
		if err != nil {
			panic("Invalid Dev")
		}
		return DecToBase(i, newBase)
	default:
		f := BaseToDec(value, base)
		return DecToBase(f, newBase)
	}
}
