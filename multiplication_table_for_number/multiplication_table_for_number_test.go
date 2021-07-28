package multiplication_table_for_number

import "testing"

func TestMultiTable(t *testing.T) {
	want := "1 * 5 = 5\n2 * 5 = 10\n3 * 5 = 15\n4 * 5 = 20\n5 * 5 = 25\n6 * 5 = 30\n7 * 5 = 35\n8 * 5 = 40\n9 * 5 = 45\n10 * 5 = 50"
	got := MultiTable(5)

	if got != want {
		t.Errorf("MultiTable() = %q, want %q", got, want)
	}
}
