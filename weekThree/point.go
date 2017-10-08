package weekThree

import "math"

type point struct {
	x, y int
}

type points []point

// compare returns true if i.y < j.y using x
// cordinates to break ties.
func (p points) compare(i, j int) bool {
	if p[i].y == p[j].y {
		return p[i].x < p[j].x
	}
	return p[i].y < p[j].y
}

func (p points) slope(i, j int) float64 {
	if p[i].y == p[j].y {
		if p[i].x == p[j].x {
			// degenerate case, equal points
			return math.Inf(-1)
		}
		// slope of horizontal line segment
		return 0
	}

	if p[i].x == p[j].x {
		// slope of vertical line segment
		return math.Inf(0)
	}
	return float64(p[j].y-p[i].y) / float64(p[j].x-p[i].x)
}
