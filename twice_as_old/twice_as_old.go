package twice_as_old

// Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).
// Expect( TwiceAsOld(36,7)).To(Equal(22))
// Expect( TwiceAsOld(55,30)).To(Equal(5))
// Expect( TwiceAsOld(42,21)).To(Equal(0))
// Expect( TwiceAsOld(22,1)).To(Equal(20))
// Expect( TwiceAsOld(29,0)).To(Equal(29))

// my solution
func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	twice := dadYearsOld - sonYearsOld*2

	if twice < 0 {
		return twice * (-1)
	}

	return twice
}

// best practise
// func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
// 	return int(math.Abs(float64(dadYearsOld - 2*sonYearsOld)))
// }
