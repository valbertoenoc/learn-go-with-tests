package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, got %v", want, got)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("expected %v, given %v, got %v", want, numbers, got)
		}
	})

	t.Run("sum of collection of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{2, 8})
		want := []int{6, 10}
		checkSums(t, got, want)
	})

	t.Run("sum of tail of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 8})
		want := []int{5, 8}
		checkSums(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 8})
		want := []int{0, 8}
		checkSums(t, got, want)
	})
}
