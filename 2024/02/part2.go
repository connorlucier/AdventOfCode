package main

func part2(values [][]int) int {
	result := 0
	for _, row := range values {
		variants := [][]int{}
		variants = append(variants, row)
		for i := range row {
			v := []int{}
			v = append(v, row[:i]...)
			v = append(v, row[i+1:]...)
			variants = append(variants, v)
		}
		if part1(variants) > 0 {
			result++
		}
	}

	return result
}
