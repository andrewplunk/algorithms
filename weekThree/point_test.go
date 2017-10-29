package weekThree

import (
	"math"
	"testing"
)

type test func(*testing.T, points)
type expectation func(int, int, bool) test

func expectSlope(i, j int, expected float64) test {
	return func(t *testing.T, p points) {
		actual := p.slope(i, j)
		if actual != expected {
			t.Fatalf("Expected %f != Actual %f", expected, actual)
		}
	}
}

func expectCompare(i, j int, expected bool) test {
	return func(t *testing.T, p points) {
		actual := p.compare(i, j)
		if actual != expected {
			t.Fatalf("Expected %t != Actual %t", expected, actual)
		}
	}
}

func expectSlopeOrder(i, j, k, expected int) test {
	return func(t *testing.T, p points) {
		actual := p.slopeOrder(i)(j, k)
		if actual != expected {
			t.Fatalf("Expected %d != Actual %d", expected, actual)
		}
	}
}

func TestPoints(t *testing.T) {
	points := points{
		point{0, 0},
		point{-1, 0},
		point{10, 10},
		point{10, 12},
	}
	tests := []test{
		// x break ties false
		expectCompare(0, 1, false),
		// x break ties true
		expectCompare(1, 0, true),
		// i.y < j.y
		expectCompare(1, 2, true),
		// j.y < i.y
		expectCompare(2, 1, false),
		// degenerate case(equal points) == negative infinity
		expectSlope(1, 1, math.Inf(-2)),
		// vertical line segment positive infinity
		expectSlope(2, 3, math.Inf(1)),
		// horizontal line segment
		expectSlope(0, 1, 0),
		// normal slope
		expectSlope(0, 2, float64(1)),
		// 0 - 0 / -1 - 0 == 0
		// 10 - 0 / 10 - 0 == 1
		expectSlopeOrder(0, 1, 2, -1),
		expectSlopeOrder(1, 2, 0, 1),
		expectSlopeOrder(1, 0, 0, 0),
		expectSlopeOrder(0, 0, 0, 0),
	}

	for _, test := range tests {
		test(t, points)
	}
}
