package greeting

import (
	"testing"
	"fmt"
)

func TestGreet(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"Bob", "Hello, Bob."},
		{"", "Hello, my friend."},
		{"JERRY", "HELLO JERRY!"},
	}
	for _, test := range tests {
		testInput := fmt.Sprintf("%v", test.input)
		t.Run(testInput, func(t *testing.T) {
			if got := Greet(test.input); got != test.want {
				t.Fatalf("greet(%v) returned %v; want %v", test.input, got, test.want)
			}
		})
	}
}
