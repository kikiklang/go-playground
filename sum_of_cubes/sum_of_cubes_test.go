package sum_of_cubes

import "testing"

func TestSumCubes1(t *testing.T) {
	want := 0
	got := SumCubes(0)

	if got != want {
		t.Errorf("SumCubes() = %q, want %q", got, want)
	}
}

func TestSumCubes2(t *testing.T) {
	want := 100
	got := SumCubes(4)

	if got != want {
		t.Errorf("SumCubes() = %q, want %q", got, want)
	}
}

func TestSumCubes3(t *testing.T) {
	want := 58155876
	got := SumCubes(123)

	if got != want {
		t.Errorf("SumCubes() = %q, want %q", got, want)
	}
}
