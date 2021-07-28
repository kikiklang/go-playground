package twice_as_old

import "testing"

func TestTwiceAsOld(t *testing.T) {
	want := 20
	got := TwiceAsOld(40, 10)

	if got != want {
		t.Errorf("TwiceAsOld() = %q, want %q", got, want)
	}
}
