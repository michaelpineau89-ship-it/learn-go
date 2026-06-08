package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "ascii", in: "hello", want: "olleh"},
		{name: "empty", in: "", want: ""},
		{name: "unicode", in: "Go🐹", want: "🐹oG"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Reverse(tc.in); got != tc.want {
				t.Fatalf("Reverse(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
