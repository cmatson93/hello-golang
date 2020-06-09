package arrays

// SumArray takes in an array of integers and returns the sum
func SumArray(numbers [3]int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count

}
