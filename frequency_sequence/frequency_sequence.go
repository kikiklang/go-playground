package frequency_sequence

import (
	"strconv"
	"strings"
)

// Return an output string that translates an input string s/$s
// by replacing each character in s/$s with a number representing the number of times
// that character occurs in s/$s and separating each number with the character(s) sep/$sep.
// freq_seq("hello world", "-"); // => "1-1-3-3-2-1-1-2-1-3-1"

func FreqSeq(str string, sep string) string {
	var stringsSlice []string
	split := strings.Split(str, "")

	for _, char := range split {
		intToString := strconv.Itoa(strings.Count(str, char))
		stringsSlice = append(stringsSlice, intToString)
	}

	return strings.Join(stringsSlice, sep)
}
