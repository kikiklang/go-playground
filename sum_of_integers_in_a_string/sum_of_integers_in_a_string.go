package sum_of_integers_in_a_string

// Your task in this kata is to implement a function that calculates the sum of the integers inside a string. For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.

import (
	"regexp"
	"strconv"
)

// func SumOfIntegersInString(strng string) int {
// 	re := regexp.MustCompile("[0-9]+")
// 	stringNumberArray := re.FindAllString(strng, -1)
// 	intArray := stringToIntArray(stringNumberArray)
// 	result := 0

// 	for i := 0; i < len(intArray); i++ {
// 		result += intArray[i]
// 	}

// 	return result
// }

// func stringToIntArray(stringArray []string) []int {
// 	var intArray = []int{}

// 	for _, i := range stringArray {
// 		j, err := strconv.Atoi(i)

// 		if err != nil {
// 			panic(err)
// 		}

// 		intArray = append(intArray, j)
// 	}

// 	return intArray
// }

func SumOfIntegersInString(strng string) int {
	re := regexp.MustCompile("[0-9]+")
	sum := 0

	for _, i := range re.FindAllString(strng, -1) {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum
}
