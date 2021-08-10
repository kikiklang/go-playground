package simple_beads_count

import "testing"

func TestCountRedBeads1(t *testing.T) {
	want := 8
	got := CountRedBeads(5)

	if got != want {
		t.Errorf("CountRedBeads() = %q, want %q", got, want)
	}
}

func TestCountRedBeads2(t *testing.T) {
	want := 0
	got := CountRedBeads(0)

	if got != want {
		t.Errorf("CountRedBeads() = %q, want %q", got, want)
	}
}

func TestCountRedBeads3(t *testing.T) {
	want := 0
	got := CountRedBeads(1)

	if got != want {
		t.Errorf("CountRedBeads() = %q, want %q", got, want)
	}
}
