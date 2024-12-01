package main

import (
	"math"
	"slices"
)

func part1(l []int, r []int) int {
	slices.Sort(l)
	slices.Sort(r)

	result := 0
	for i, _ := range l {
		result += int(math.Abs(float64(l[i] - r[i])))
	}

	return result
}
