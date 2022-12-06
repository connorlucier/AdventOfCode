package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	stacks := make(map[int]string, 9)

	parsedState := false
	for scanner.Scan() {
		line := scanner.Text()

		if !parsedState {
			for i, c := range line {
				if c < 'A' || c > 'Z' {
					continue
				}

				stacks[i/4+1] += string(c)
			}

			parsedState = parsedState || line == ""
			continue
		}

		var crates, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &crates, &from, &to)
		stacks[to] = rev(stacks[from][:crates]) + stacks[to]
		stacks[from] = stacks[from][crates:]
	}

	for key := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		if len(stacks[key]) > 0 {
			fmt.Print(stacks[key][:1])
		}
	}
	fmt.Println()
}

func rev(input string) string {
	var bytes []byte
	for i := len(input) - 1; i >= 0; i-- {
		bytes = append(bytes, input[i])
	}

	return string(bytes)
}
