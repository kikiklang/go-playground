package consecutive_letters

import "testing"

func TestConsecutiveLetters1(t *testing.T) {
	want := true
	got := ConsecutiveLetters("abc")

	if got != want {
		t.Errorf("ConsecutiveLetters() = %t, want %t", got, want)
	}
}

func TestConsecutiveLetters2(t *testing.T) {
	want := true
	got := ConsecutiveLetters("dabc")

	if got != want {
		t.Errorf("ConsecutiveLetters() = %t, want %t", got, want)
	}
}

func TestConsecutiveLetters3(t *testing.T) {
	want := true
	got := ConsecutiveLetters("z")

	if got != want {
		t.Errorf("ConsecutiveLetters() = %t, want %t", got, want)
	}
}

func TestConsecutiveLetters4(t *testing.T) {
	want := false
	got := ConsecutiveLetters("abd")

	if got != want {
		t.Errorf("ConsecutiveLetters() = %t, want %t", got, want)
	}
}

func TestConsecutiveLetters5(t *testing.T) {
	want := false
	got := ConsecutiveLetters("abbc")

	if got != want {
		t.Errorf("ConsecutiveLetters() = %t, want %t", got, want)
	}
}
