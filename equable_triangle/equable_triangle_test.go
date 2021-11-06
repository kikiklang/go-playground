package equable_triangle

import "testing"

func TestEquableTriangle1(t *testing.T) {
	want := true
	got := EquableTriangle(5, 12, 13)

	if got != want {
		t.Errorf("EquableTriangle() = %t, want %t", got, want)
	}
}

func TestEquableTriangle2(t *testing.T) {
	want := false
	got := EquableTriangle(2, 3, 4)

	if got != want {
		t.Errorf("EquableTriangle() = %t, want %t", got, want)
	}
}
