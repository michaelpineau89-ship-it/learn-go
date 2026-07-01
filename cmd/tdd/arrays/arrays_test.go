package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("expected '%v' got '%v'", want, got)
		}
	})
}

func TestAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("expected '%v' got '%v'", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
	t.Run("Make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

	t.Run("Attempt to sum an empty Slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSum(t, got, want)

	})
}
