package over_the_road

// You've just moved into a perfectly straight street with exactly n identical houses on either side of the road.
// Naturally, you would like to find out the house number of the people on the other side of the street. The street looks something like this:

func OverTheRoad(address int, streetLength int) int {
	return (streetLength*2 - address) + 1
}
