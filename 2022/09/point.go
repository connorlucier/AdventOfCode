package main

import "math"

type point struct {
	X int
	Y int
}

func (p1 point) DistanceTo(p2 point) int {
	diffX := math.Abs(float64(p2.X - p1.X))
	diffY := math.Abs(float64(p2.Y - p1.Y))
	return int(math.Sqrt(diffX*diffX + diffY*diffY))
}

func Contains(pts []point, val point) bool {
	for _, pt := range pts {
		if pt.X == val.X && pt.Y == val.Y {
			return true
		}
	}
	return false
}
