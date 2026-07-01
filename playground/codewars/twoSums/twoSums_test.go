package main

import "testing"

func TestTwoSums(t *testing.T) {
	got := TwoSum([]int{1, 2, 3}, 4)
	want := [2]int{0, 2}

	if got != want {
		t.Errorf("want: '%v' got: '%v'", want, got)
	}

}
