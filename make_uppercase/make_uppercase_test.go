package make_uppercase

import "testing"

func TestMakeUpperCase1(t *testing.T) {
	want := "HELLO WORLD !"
	got := MakeUpperCase("hello world !")

	if got != want {
		t.Errorf("MakeUpperCase() = %q, want %q", got, want)
	}
}

func TestMakeUpperCase2(t *testing.T) {
	want := "1,2,3 HELLO WORLD!"
	got := MakeUpperCase("1,2,3 hellO World!")

	if got != want {
		t.Errorf("MakeUpperCase() = %q, want %q", got, want)
	}
}
