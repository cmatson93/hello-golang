package iterations

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got %q but expected %q", got, expected)
		}
	}

	t.Run("returns c repeated 5 times", func(t *testing.T) {
		got := Repeat("c", 5)
		expected := "ccccc"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("returns c repeated x ammount of times specified", func(t *testing.T) {
		got := Repeat("c", 2)
		expected := "cc"
		assertCorrectMessage(t, got, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
