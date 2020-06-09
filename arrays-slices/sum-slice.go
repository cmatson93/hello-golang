package arrays

// SumSlice takes in an array of x integers and returns the sum
func SumSlice(numbers []int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count

}
