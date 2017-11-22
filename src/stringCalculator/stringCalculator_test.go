package stringCalculator

import (
	"testing"
)

func TestStringCalculator(t *testing.T) {
	var tests = []struct {
		input string
		want  interface{}
	}{
		{"", 0},
		{"1", 1},
		{"1,1", 2},
		{"2\n3", 5},
		{"1,1,1,2,2,3", 10},
		{"//t", 0},
		{"//t1", 1},
		{"//p\n34535345p343535345p34234234p-178684535", -1},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if got := Add(test.input); got != test.want {
				t.Fatalf("StringCalculator(%v) returned %v, want %v", test.input, got, test.want)
			}
		})
	}
}
