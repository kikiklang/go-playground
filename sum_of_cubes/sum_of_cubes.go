package sum_of_cubes

// Write a function that takes a positive integer n, sums all the cubed values from 1 to n, and returns that sum.
// Assume that the input n will always be a positive integer.

// my solution
func SumCubes(n int) int {
	result := 0

	for i := 0; i <= n; i++ {
		result += i * i * i
	}

	return result
}

// best practise

// func SumCubes(n int) int {
// 	n = n * (n + 1)
// 	return n * n / 4
// }
