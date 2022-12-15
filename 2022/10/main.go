package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := parseInput()
	part1(input)
	part2(input)
}

func parseInput() []instruction {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var input []instruction
	for scanner.Scan() {
		line := scanner.Text()

		var command string
		var val int
		fmt.Sscanf(line, "%s %d", &command, &val)

		input = append(input, instruction{})

		if command == "addx" {
			input = append(input, instruction{
				Name:  "addx",
				Value: val,
			})
		}
	}

	return input
}
