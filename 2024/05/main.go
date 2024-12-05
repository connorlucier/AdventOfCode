package main

import (
	"bufio"
	"fmt"
	"os"
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
