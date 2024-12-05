package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	updates, rules := parseInput()
	fmt.Println(part1(updates, rules), part2(updates, rules))
}

func parseInput() ([][]int, map[int][]int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	isRule := true
	rules := map[int][]int{}
	updates := [][]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			isRule = false
			continue
		}

		if isRule {
			strs := strings.Split(text, "|")
			l, _ := strconv.Atoi(strs[0])
			r, _ := strconv.Atoi(strs[1])

			if rules[l] == nil {
				rules[l] = []int{r}
			} else {
				rules[l] = append(rules[l], r)
			}
		} else {
			pages := strings.Split(text, ",")
			up := make([]int, len(pages))
			for i, p := range pages {
				num, _ := strconv.Atoi(p)
				up[i] = num
			}
			updates = append(updates, up)
		}
	}
	return updates, rules
}

func validateList(list []int, rules map[int][]int) (bool, []int) {
	isValid := true
	sortFn := func(a int, b int) int {
		if rules[a] != nil && slices.Contains(rules[a], b) {
			return -1
		} else if rules[b] != nil && slices.Contains(rules[b], a) {
			return 1
		}
		return 0
	}
	for i, v := range slices.Backward(list) {
		if rules[v] != nil {
			for _, u := range list[:i] {
				if slices.Contains(rules[v], u) {
					isValid = false
					break
				}
			}
		}
	}
	sorted := make([]int, len(list))
	copy(sorted, list)
	slices.SortFunc(sorted, sortFn)

	return isValid, sorted
}
