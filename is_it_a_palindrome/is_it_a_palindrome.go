package is_it_a_palindrome

import (
	"strings"
)

// Write function that checks if a given string (case insensitive) is a palindrome.
func IsPalindrome(str string) bool {
	var runes = []rune(strings.ToLower(str))

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) == strings.ToLower(str)
}
