package arrays

// SumAll takes in slices and returns a slice containing all totals from each slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, slice := range numbersToSum {
		sliceSum := SumSlice(slice)
		sums = append(sums, sliceSum)
	}
	return
}
