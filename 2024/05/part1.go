package main

func part1(updates [][]int, rules map[int][]int) int {
	result := 0
	for _, up := range updates {
		isValid, _ := validateList(up, rules)
		if isValid {
			result += up[len(up)/2]
		}
	}
	return result
}
