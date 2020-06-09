package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Christina")
	want := "Hello, Christina"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
