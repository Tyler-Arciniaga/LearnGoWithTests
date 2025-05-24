package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to a specific person", func(t *testing.T) {
		got := Hello("Tyler")
		want := "Hello Tyler"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	got := Hello("Tyler")
	want := "Hello Tyler"
	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
