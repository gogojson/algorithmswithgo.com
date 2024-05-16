package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	result := ""
	hexL := "0123456789ABCDEF"

	for {
		result = string(hexL[dec%base]) + result
		dec = dec / base
		if dec == 0 {
			break
		}
	}
	return result
}
