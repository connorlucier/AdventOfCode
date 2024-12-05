package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, xLocs, aLocs := parseInput()
	fmt.Println(part1(grid, xLocs), part2(grid, aLocs))
}

func parseInput() ([][]rune, []Coord, []Coord) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	grid := [][]rune{}
	xLocs := []Coord{}
	aLocs := []Coord{}
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, []rune(text))
		for j, c := range text {
			coord := Coord{
				r: i,
				c: j,
			}
			if c == 'X' {
				xLocs = append(xLocs, coord)
			} else if c == 'A' {
				aLocs = append(aLocs, coord)
			}
		}
		i++
	}
	return grid, xLocs, aLocs
}

type Coord struct {
	r int
	c int
}
