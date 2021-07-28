package quarter

// Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

// my solution
// func QuarterOf(month int) int {
// 	if month >= 1 && month < 4 {
// 		return 1
// 	} else if month >= 4 && month < 7 {
// 		return 2
// 	} else if month >= 7 && month < 10 {
// 		return 3
// 	} else if month >= 10 && month < 13 {
// 		return 4
// 	}

// 	return 0
// }

// best practise
func QuarterOf(month int) (result int) {

	switch month {
	case 1, 2, 3:
		result = 1
	case 4, 5, 6:
		result = 2
	case 7, 8, 9:
		result = 3
	case 10, 11, 12:
		result = 4
	}
	return
}
