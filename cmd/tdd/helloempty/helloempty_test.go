package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying Hello to People", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertRightMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertRightMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertRightMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mathieu", "French")
		want := "Bonjour, Mathieu"

		assertRightMessage(t, got, want)
	})

	t.Run("in Esperanto", func(t *testing.T) {
		got := Hello("Marko", "Esperanto")
		want := "Saluton, Marko"

		assertRightMessage(t, got, want)
	})
}

func assertRightMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want %q", got, want)
	}
}
