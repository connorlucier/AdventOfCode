package main

func part1(values [][]int) int {
	result := 0
	for _, row := range values {
		if isRowSafe(row) {
			result++
		}
	}
	return result
}
