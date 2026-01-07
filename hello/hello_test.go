package main

import (
	"testing"
)


func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Koki", "English")
		want := "Hello, Koki"

		assertEqual(t, got, want)
	})


	t.Run("saying hello to nobody", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertEqual(t, got, want)

	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Koki", "Spanish")
		want := "Hola, Koki"

		assertEqual(t, got, want)

	})
}


func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}


