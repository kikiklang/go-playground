package maximum_length_difference

func MxDifLg(a1 []string, a2 []string) int {

	if len(a1) <= 0 || len(a2) <= 0 {
		return -1
	}

	a1Min := len(a1[0])
	a1Max := len(a1[0])
	a2Min := len(a2[0])
	a2Max := len(a2[0])
	var result int

	for _, s := range a1 {
		strLength := len(s)

		if strLength < a1Min {
			a1Min = strLength
		}

		if strLength > a1Max {
			a1Max = strLength
		}
	}

	for _, s := range a2 {
		strLength := len(s)

		if strLength < a2Min {
			a2Min = strLength
		}

		if strLength > a2Max {
			a2Max = strLength
		}
	}

	if a1Max-a2Min >= a2Max-a1Min {
		result = a1Max - a2Min
	} else {
		result = a2Max - a1Min
	}

	return result
}
