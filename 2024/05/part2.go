package main

func part2(updates [][]int, rules map[int][]int) int {
	result := 0
	for _, up := range updates {
		isValid, sorted := validateList(up, rules)
		if !isValid {
			result += sorted[len(sorted)/2]
		}
	}
	return result
}
