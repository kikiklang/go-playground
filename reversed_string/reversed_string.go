package reversed_string

import "fmt"

// Complete the solution so that it reverses the string passed into it.

// my solution
func ReversedString(word string) string {
	var rns = []rune(word)
	fmt.Println(rns)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
