package fix_string_case

import "testing"

func TestFixStringCase1(t *testing.T) {
	want := "code"
	got := FixStringCase("coDe")

	if got != want {
		t.Errorf("MakeUpperCase() = %q, want %q", got, want)
	}
}

func TestFixStringCase2(t *testing.T) {
	want := "code"
	got := FixStringCase("CoDe")

	if got != want {
		t.Errorf("MakeUpperCase() = %q, want %q", got, want)
	}
}

func TestFixStringCase3(t *testing.T) {
	want := "CODE"
	got := FixStringCase("CoDE")

	if got != want {
		t.Errorf("MakeUpperCase() = %q, want %q", got, want)
	}
}
