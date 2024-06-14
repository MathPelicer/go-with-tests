package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"
	assertCorrectMessage(t, expected, result)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func assertCorrectMessage(t testing.TB, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}
