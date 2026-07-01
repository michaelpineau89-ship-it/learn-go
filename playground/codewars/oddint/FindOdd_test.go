package main

import "testing"

func TestFindOdd(t *testing.T) {
	t.Run("testing odd numbers - 1", func(t *testing.T) {
		got := FindOdd([]int{1, 2, 1})
		want := 2

		if got != want {
			t.Errorf("wanted '%v', got '%v' ", want, got)
		}
	})

	t.Run("second Test", func(t *testing.T) {
		got := FindOdd([]int{1, 1, 2, 1, 2})
		want := 1

		if got != want {
			t.Errorf("wanted '%v', got '%v' ", want, got)
		}
	})

}
