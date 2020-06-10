package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, expected int, numbers [3]int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d but expected %d given %v", got, expected, numbers)
		}
	}

	t.Run("returns 15 when given an array of 3 fives", func(t *testing.T) {
		numbers := [3]int{5, 5, 5}
		got := SumArray(numbers)
		expected := 15
		assertCorrectMessage(t, got, expected, numbers)
	})

	t.Run("returns sum when given an array of any items", func(t *testing.T) {
		numbers := []int{5, 5, 5, 5, 5}
		got := SumSlice(numbers)
		expected := 25
		if got != expected {
			t.Errorf("got %d but expected %d given %v", got, expected, numbers)
		}
	})

	t.Run("takes a varying number of slices and returns new slice with each total", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
