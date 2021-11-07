package functional_addition

// Create a function add(n)/Add(n) which returns a function that always adds n to any number

func Add(i int) func(int) int {
	return func(x int) int {
		return i + x
	}
}
