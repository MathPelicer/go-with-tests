package arraynslices

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("5 number collection", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if expected != result {
			t.Errorf("Expected %d, but got %d, %v", expected, result, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(expected, result) {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, result []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %d, but got %d", expected, result)
		}

	}
	t.Run("sum some slices", func(t *testing.T) {
		expected := []int{2, 9}
		result := SumAllTails([]int{1, 2}, []int{0, 9})
		checkSums(t, expected, result)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		expected := []int{0, 9}
		result := SumAllTails([]int{}, []int{0, 9})
		checkSums(t, expected, result)
	})
}
