package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

func RecursionSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + RecursionSum(numbers[1:])
}
