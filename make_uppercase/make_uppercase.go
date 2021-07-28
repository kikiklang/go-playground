package make_uppercase

import "strings"

// Write a function which converts the input string to uppercase.
// Expect(MakeUpperCase("hello")).To(Equal("HELLO"))
// Expect(MakeUpperCase("hello world")).To(Equal("HELLO WORLD"))
// Expect(MakeUpperCase("hello world !")).To(Equal("HELLO WORLD !"))
// Expect(MakeUpperCase("heLlO wORLd !")).To(Equal("HELLO WORLD !"))
// Expect(MakeUpperCase("1,2,3 hello world!")).To(Equal("1,2,3 HELLO WORLD!"))

// my solution
func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}

// best practise
// func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
// 	return int(math.Abs(float64(dadYearsOld - 2*sonYearsOld)))
// }
