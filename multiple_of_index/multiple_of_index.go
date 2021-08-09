package multiple_of_index

// Return a new array consisting of elements which are multiple of their own index in input array (length > 1).

// [22, -6, 32, 82, 9, 25] => [-6, 32, 25]

func MultipleOfIndex(ints []int) []int {
	multiples := []int{}

	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			multiples = append(multiples, ints[i])
		}
	}

	return multiples
}
