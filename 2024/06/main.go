package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, start := parseInput()
	result, corners := traverse(grid, start)
	fmt.Println(part1(result), part2(grid, start, corners))
}

func parseInput() ([][]rune, Coord) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	grid := [][]rune{}
	var start Coord
	for scanner.Scan() {
		text := scanner.Text()
		row := []rune(text)
		for c, v := range row {
			if v == Up {
				start = Coord{
					row: len(grid),
					col: c,
				}
			}
		}
		grid = append(grid, row)
	}
	return grid, start
}

func traverse(grid [][]rune, start Coord) ([][]rune, []Coord) {
	r := start.row
	c := start.col
	path := []Coord{}
	for {
		dir, newR, newC := move(grid, r, c)
		if newR != r || newC != c {
			path = append(path, Coord{
				row: r,
				col: c,
			})
		}
		if isOutOfBounds(grid, newR, newC) {
			break
		}
		grid[newR][newC] = dir
		r = newR
		c = newC
	}

	return grid, path
}

func isOutOfBounds(grid [][]rune, r int, c int) bool {
	return r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0])
}

func move(grid [][]rune, r int, c int) (rune, int, int) {
	dir := grid[r][c]
	newR, newC := next(dir, r, c)

	if !isOutOfBounds(grid, newR, newC) && grid[newR][newC] == Object {
		return turnRight(dir), r, c
	}
	return dir, newR, newC
}

func next(dir rune, r int, c int) (int, int) {
	switch dir {
	case Up:
		return r - 1, c
	case Right:
		return r, c + 1
	case Down:
		return r + 1, c
	case Left:
		return r, c - 1
	default:
		return r, c
	}
}

func turnRight(dir rune) rune {
	switch dir {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		return Up
	}
}

func countVisited(grid [][]rune) int {
	result := 0
	for _, row := range grid {
		for _, val := range row {
			if val != Empty && val != Object {
				result++
			}
		}
	}
	return result
}

func verticalSlice(arr [][]rune, c int, startR int, endR int) []rune {
	result := []rune{}
	for r := startR; r < endR; r++ {
		result = append(result, arr[r][c])
	}
	return result
}

type Coord struct {
	row int
	col int
}

const (
	Up     = '^'
	Right  = '>'
	Down   = 'v'
	Left   = '<'
	Object = '#'
	Empty  = '.'
)
