package iterations

var repeatCount = 5

// Repeat is a function that takes in a value and repeats it 5 times
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
