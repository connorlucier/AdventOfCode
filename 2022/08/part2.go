package main

import "fmt"

func part2(grid [][]tree) {
	maxScore := 0
	for _, row := range grid[1 : len(grid)-1] {
		for _, tree := range row[1 : len(row)-1] {
			score := scenicScore(tree, grid)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)
}

func scenicScore(t tree, grid [][]tree) int {
	up := 0
	for row := t.Row - 1; row >= 0; row-- {
		up++
		if grid[row][t.Col].Height >= t.Height {
			break
		}
	}

	down := 0
	for row := t.Row + 1; row < len(grid); row++ {
		down++
		if grid[row][t.Col].Height >= t.Height {
			break
		}
	}

	left := 0
	for col := t.Col - 1; col >= 0; col-- {
		left++
		if grid[t.Row][col].Height >= t.Height {
			break
		}
	}

	right := 0
	for col := t.Col + 1; col < len(grid); col++ {
		right++
		if grid[t.Row][col].Height >= t.Height {
			break
		}
	}

	return up * down * left * right
}
