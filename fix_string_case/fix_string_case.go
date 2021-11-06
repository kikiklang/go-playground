package fix_string_case

import (
	"strings"
	"unicode"
)

// In this Kata, you will be given a string that may have mixed uppercase and lowercase letters
// and your task is to convert that string to either lowercase only or uppercase only based on:
// make as few changes as possible.
// if the string contains equal number of uppercase and lowercase letters, convert the string to lowercase.

func FixStringCase(str string) string {
	lowercaseCount := 0
	uppercaseCount := 0

	for _, char := range str {
		if unicode.IsLower(char) {
			lowercaseCount++
		} else {
			uppercaseCount++
		}
	}

	if lowercaseCount >= uppercaseCount {
		return strings.ToLower(str)
	} else {
		return strings.ToUpper(str)
	}
}
