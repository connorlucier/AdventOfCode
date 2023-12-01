package main

import (
	"bufio"
	"os"
)

func main() {
	grid, start, end := parseInput()
	part1(grid, start, end)
	part2(grid, start, end)
}

func parseInput() (grid [][]*point, start *point, end *point) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []*point{})

		for col, char := range line {
			var height int
			if char == 'S' {
				height = 0
			} else if char == 'E' {
				height = 25
			} else {
				height = int(char - 97)
			}

			grid[row] = append(grid[row], &point{
				X:      row,
				Y:      col,
				Height: height,
			})

			if char == 'S' {
				start = grid[row][col]
			}
			if char == 'E' {
				end = grid[row][col]
			}
		}
	}

	return grid, start, end
}
