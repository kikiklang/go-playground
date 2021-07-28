package maximum_length_difference

import "testing"

func TestMxDifLgr1(t *testing.T) {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	want := 13
	got := MxDifLg(a1, a2)

	if got != want {
		t.Errorf("MxDifLg() = %q, want %q", got, want)
	}
}

func TestMxDifLgr2(t *testing.T) {
	a1 := []string{}
	a2 := []string{"cccooommaaqqoxii"}
	want := -1
	got := MxDifLg(a1, a2)

	if got != want {
		t.Errorf("MxDifLg() = %q, want %q", got, want)
	}
}
