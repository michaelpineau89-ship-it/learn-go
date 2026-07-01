package main

import (
	"testing"
)

func TestWeirdStrings(t *testing.T) {

	t.Run("basic_test", func(t *testing.T) {
		want := "ThIsIsAsTrInG"
		got := toWeirdCase("thisisastring")

		if want != got {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("basic_test", func(t *testing.T) {
		want := "AbC DeF"
		got := toWeirdCase("abc def")

		if want != got {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("basic_test", func(t *testing.T) {
		want := "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"
		got := toWeirdCase("This is a test Looks like you passed")

		if want != got {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

}
