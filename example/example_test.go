package GolangKatas

import "testing"

func TestExampleFunction(t *testing.T) {
	input := 1
	want := 1
	got := ExampleFunction(input)
	if got != want{
		t.Errorf("ExampleFunction(%v)", input)
		t.Errorf("returned %v, want %v", got, want)
	}
}
