package consecutive_letters

import (
	"sort"
	"strings"
)

// Rules are: (1) the letters are adjacent in the English alphabet, and (2) each letter occurs only once.
// For example:
// solve("abc") = True, because it contains a,b,c
// solve("abd") = False, because a, b, d are not consecutive/adjacent in the alphabet, and c is missing.
// solve("dabc") = True, because it contains a, b, c, d
// solve("abbc") = False, because b does not occur once.
// solve("v") = True

// func ConsecutiveLetters(s string) bool {
// 	if checkDuplicate(s) {
// 		return false
// 	}

// 	if !checkConsecutive(s) {
// 		return false
// 	}

// 	return true
// }

// func checkDuplicate(w string) bool {
// 	chars := []rune(w)
// 	le := len(chars)

// 	if le > 1 {
// 		for i := 0; i < le-1; i++ {
// 			for j := i + 1; j < le; j++ {
// 				if chars[j] == chars[i] {
// 					return true
// 				}
// 			}
// 		}
// 	}

// 	return false
// }

// func sortString(w string) string {
// 	s := strings.Split(w, "")
// 	sort.Strings(s)

// 	return strings.Join(s, "")
// }

// func checkConsecutive(w string) bool {
// 	sortedString := sortString(w)
// 	chars := []rune(sortedString)
// 	le := len(chars)

// 	for i := 0; i < le-1; i++ {
// 		if chars[i+1]-chars[i] != 1 {
// 			return false
// 		}
// 	}

// 	return true
// }

/////////////////////////////////////////////////////////////
//best practise
func ConsecutiveLetters(s string) bool {
	split := strings.Split(s, "")
	sort.Strings(split)
	sorted := strings.Join(split, "")

	return strings.Contains("abcdefghijklmnopqrstuvwxyz", sorted)
}
