package main

import "slices"

func part1(updates [][]int, rules map[int][]int) int {
	result := 0
	for _, up := range updates {
		if isValidOrder(up, rules) {
			result += up[len(up)/2]
		}
	}
	return result
}

func isValidOrder(up []int, rules map[int][]int) bool {
	for i, v := range slices.Backward(up) {
		if rules[v] != nil {
			left := up[:i]
			for _, l := range left {
				if slices.Contains(rules[v], l) {
					return false
				}
			}
		}
	}
	return true
}
