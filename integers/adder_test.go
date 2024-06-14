package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	assertCorrectMessage(t, expected, sum)
}

func ExampleAdd() {
	sum := Add(5, 4)
	fmt.Println(sum)
	// Output: 9
}

func assertCorrectMessage(t testing.TB, expected, result int) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected '%d', but got '%d'", expected, result)
	}
}
