package equable_triangle

import "math"

// A triangle is called an equable triangle if its area equals its perimeter.
// Return true, if it is an equable triangle, else return false.
// You will be provided with the length of sides of the triangle. Happy Coding!

func EquableTriangle(a, b, c int) bool {
	s1 := float64(a)
	s2 := float64(b)
	s3 := float64(c)
	p := (s1 + s2 + s3) / 2
	area := math.Sqrt(p*((p-s1)*(p-s2)*(p-s3))) / 2

	return p == area
}
