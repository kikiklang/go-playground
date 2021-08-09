package speed_control

import "math"

func Gps(s int, segments []float64) int {
	var maxSpeed float64

	for index := 1; index < len(segments); index++ {
		startSegment := segments[index-1]
		endSegment := segments[index]
		kmPerHour := 3600.0 * (endSegment - startSegment) / float64(s)
		maxSpeed = math.Max(maxSpeed, kmPerHour)
	}

	return int(maxSpeed)
}
