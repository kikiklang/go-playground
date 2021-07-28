package quarter

import "testing"

func TestQuarter(t *testing.T) {
	want := 2
	got := QuarterOf(5)

	if got != want {
		t.Errorf("QuarterOf() = %q, want %q", got, want)
	}
}
