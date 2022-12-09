package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() [][]tree {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var grid [][]tree

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]tree, len(line)))

		for col, num := range line {
			row := len(grid) - 1
			val, _ := strconv.Atoi(string(num))
			grid[row][col] = tree{
				Height: val,
				Row:    row,
				Col:    col,
			}
		}
	}

	totalVisible := len(grid)*2 + (len(grid[0])-2)*2

	for _, row := range grid[1 : len(grid)-1] {
		for _, tree := range row[1 : len(row)-1] {
			if isVisible(tree, grid) {
				totalVisible++
			}
		}
	}

	fmt.Println(totalVisible)
	return grid
}

func isVisible(t tree, grid [][]tree) bool {
	up := true
	for row := t.Row - 1; row >= 0; row-- {
		if grid[row][t.Col].Height >= t.Height {
			up = false
			break
		}
	}

	down := true
	for row := t.Row + 1; row < len(grid); row++ {
		if grid[row][t.Col].Height >= t.Height {
			down = false
			break
		}
	}

	left := true
	for col := t.Col - 1; col >= 0; col-- {
		if grid[t.Row][col].Height >= t.Height {
			left = false
			break
		}
	}

	right := true
	for col := t.Col + 1; col < len(grid); col++ {
		if grid[t.Row][col].Height >= t.Height {
			right = false
			break
		}
	}

	return up || down || left || right
}
