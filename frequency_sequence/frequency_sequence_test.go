package frequency_sequence

import "testing"

func TestFreqSeq1(t *testing.T) {
	want := "1-1-3-3-2-1-1-2-1-3-1"
	got := FreqSeq("hello world", "-")

	if got != want {
		t.Errorf("FreqSeq() = %q, want %q", got, want)
	}
}

func TestFreqSeq2(t *testing.T) {
	want := "1:7:7:7:7:7:7:7"
	got := FreqSeq("19999999", ":")

	if got != want {
		t.Errorf("FreqSeq() = %q, want %q", got, want)
	}
}

func TestFreqSeq3(t *testing.T) {
	want := "3x3x3x2x2x1"
	got := FreqSeq("^^^**$", "x")

	if got != want {
		t.Errorf("FreqSeq() = %q, want %q", got, want)
	}
}
