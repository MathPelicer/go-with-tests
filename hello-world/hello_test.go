package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		result := Hello("Math", "")
		expect := "Hello, Math"

		assertCorrectMessage(t, expect, result)
	})
	t.Run("say 'Hello, world' when an emoty string is supplied", func(t *testing.T) {
		result := Hello("", "")
		expect := "Hello, world"
		assertCorrectMessage(t, expect, result)
	})
	t.Run("now in portuguese", func(t *testing.T) {
		result := Hello("Matheus", "portuguese")
		expect := "Ola, Matheus"
		assertCorrectMessage(t, expect, result)
	})
}

func assertCorrectMessage(t testing.TB, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected result: %q, but got: %q", expected, result)
	}

}
