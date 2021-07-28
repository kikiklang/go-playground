package even_or_odd

import "testing"

func TestOdd(t *testing.T) {
	want := "Odd"
	got := EvenOrOdd(5)

	if got != want {
		t.Errorf("EvenOrOdd() = %q, want %q", got, want)
	}
}

func TestEven(t *testing.T) {
	want := "Even"
	got := EvenOrOdd(36)

	if got != want {
		t.Errorf("EvenOrOdd() = %q, want %q", got, want)
	}
}
