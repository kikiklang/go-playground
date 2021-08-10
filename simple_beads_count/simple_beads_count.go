package simple_beads_count

// Two red beads are placed between every two blue beads. There are N blue beads. After looking at the arrangement below work out the number of red beads.

func CountRedBeads(n int) int {
	if n < 1 {
		return 0
	}

	return n*2 - 2
}
