package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Christina", "")
		want := "Hello, Christina"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, Sonia' when spanish is suplied as a language", func(t *testing.T) {
		got := Hello("Sonia", "Spanish")
		want := "Hola, Sonia"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Bonjor, Antoni' when french is suplied as a language", func(t *testing.T) {
		got := Hello("Antoni", "French")
		want := "Bonjor, Antoni"
		assertCorrectMessage(t, got, want)
	})
}
