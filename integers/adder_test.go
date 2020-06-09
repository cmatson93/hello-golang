package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("got %d but expected %d", sum, expected)
		}
	}

	t.Run("when supplied 1 and 1 return 2", func(t *testing.T) {
		sum := Add(1, 1)
		expected := 2
		assertCorrectMessage(t, sum, expected)
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
