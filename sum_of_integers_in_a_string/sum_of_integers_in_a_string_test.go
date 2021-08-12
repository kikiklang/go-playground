package sum_of_integers_in_a_string

import (
	"testing"
)

func TestSumOfIntegersInString1(t *testing.T) {
	want := 3868
	got := SumOfIntegersInString("The Great Depression lasted from 1929 to 1939.")

	if got != want {
		t.Errorf("SumOfIntegersInString() = %q, want %q", got, want)
	}
}

func TestSumOfIntegersInString2(t *testing.T) {
	want := 16
	got := SumOfIntegersInString("12.4")

	if got != want {
		t.Errorf("SumOfIntegersInString() = %q, want %q", got, want)
	}
}

func TestSumOfIntegersInString3(t *testing.T) {
	want := 3
	got := SumOfIntegersInString("h3ll0w0rld")

	if got != want {
		t.Errorf("SumOfIntegersInString() = %q, want %q", got, want)
	}
}

func TestSumOfIntegersInString4(t *testing.T) {
	want := 0
	got := SumOfIntegersInString("Dogs are our best friends.")

	if got != want {
		t.Errorf("SumOfIntegersInString() = %q, want %q", got, want)
	}
}

func TestSumOfIntegersInString5(t *testing.T) {
	want := 3635
	got := SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog")

	if got != want {
		t.Errorf("SumOfIntegersInString() = %q, want %q", got, want)
	}
}

// func TestStringArrayToIntArray(t *testing.T) {
// 	want := []int{17, 24, 29}
// 	got := stringToIntArray([]string{"17", "24", "29"})

// 	if len(got) != len(want) {
// 		t.Errorf("TestStringArrayToIntArray() = %q, want %q", got, want)
// 	}

// 	for index, value := range got {
// 		if value != want[index] {
// 			t.Errorf("TestStringArrayToIntArray() = %q, want %q", got, want)
// 		}
// 	}
// }
