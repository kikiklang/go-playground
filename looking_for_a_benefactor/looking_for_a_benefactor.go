package looking_for_a_benefactor

import "math"

// The accounts of the "Fat to Fit Club (FFC)" association are supervised by John as a volunteered accountant.
// The association is funded through financial donations from generous benefactors. John has a list of the first n donations: [14, 30, 5, 7, 9, 11, 15]
// He wants to know how much the next benefactor should give to the association so that the average of the first n + 1 donations should reach an average of 30.
// After doing the math he found 149. He thinks that he made a mistake. Could you help him?

func NewAvg(arr []float64, navg float64) int64 {
	var sum float64

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	result := navg*float64(len(arr)+1) - sum

	if result <= 0 {
		return -1
	}

	return int64(math.Ceil(result))
}
