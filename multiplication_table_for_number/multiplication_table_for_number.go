package multiplication_table_for_number

import (
	"strconv"
	"strings"
)

//my solution
func MultiTable(number int) string {
	var fullString string

	for i := 1; i <= 10; i++ {
		result := number * i
		str := strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(result) + "\n"
		fullString += str
	}

	return strings.TrimSuffix(fullString, "\n")
}

// best practise
// func MultiTable(number int) string {
//   superstring := ""
//   for i := 1; i < 11; i++ {
//     superstring += fmt.Sprintf("%d * %d = %d\n", i, number, number * i)
//   }
//   return strings.TrimRight(superstring, "\n")
// }
