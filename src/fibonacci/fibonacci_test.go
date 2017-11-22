package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
	}
	for _, test := range tests {
		testInput := fmt.Sprintf("%v", test.input)
		t.Run(testInput, func(t *testing.T) {
			if got := Fibonacci(test.input); got != test.want {
				t.Fatalf("Fibonacci(%v) returned %v, want %v", test.input, got, test.want)
			}
		})
	}
}

func TestSameFibonacci(t *testing.T) {
	n := 10
	got := Fibonacci(n)
	gotRec := TailRecursive(n)
	if got != gotRec {
		t.Errorf("Fibonacci tail recursive version doesn't returns the same value the regular one does")
		t.Errorf("Fibonacci(%v) returned %v", n, got)
		t.Errorf("FibonacciTailRec(%v) returned %v", n, gotRec)
	}
}

func TestFibonacciTailRec(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{20, 6765},
		{100, 3736710778780434371},
	}
	for _, test := range tests {
		testInput := fmt.Sprintf("%v", test.input)
		t.Run(testInput, func(t *testing.T) {
			if got := TailRecursive(test.input); got != test.want {
				t.Fatalf("FibonacciTailRec(%v) returned %v, want %v", test.input, got, test.want)
			}
		})
	}
}
