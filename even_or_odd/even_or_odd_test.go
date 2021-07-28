package even_or_odd

import "testing"

func TestEvenOrOdd(t *testing.T) {
	want := "Odd"
	got := EvenOrOdd(5)

	if got != want {
		t.Errorf("EvenOrOdd() = %q, want %q", got, want)
	}
}
