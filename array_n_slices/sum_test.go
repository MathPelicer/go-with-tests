package arraynslices

import "testing"

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
