package main

func part2(values [][]int) int {
	result := 0
	for _, row := range values {
		variants := getVariants(row)
		for _, v := range variants {
			if isRowSafe(v) {
				result++
				break
			}
		}
	}
	return result
}

func getVariants(values []int) [][]int {
	result := append([][]int{}, values)
	for i := range values {
		v := append([]int{}, values[:i]...)
		v = append(v, values[i+1:]...)
		result = append(result, v)
	}
	return result
}
