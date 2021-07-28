package reversed_string

import "testing"

func TestReversedString(t *testing.T) {
	want := "dlrow"
	got := ReversedString("world")

	if got != want {
		t.Errorf("ReversedString() = %q, want %q", got, want)
	}
}
