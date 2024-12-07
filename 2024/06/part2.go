package main

import (
	"slices"
)

func part2(grid [][]rune, start Coord, path []Coord) int {
	result := 0
	corners := []Coord{}
	for i := 1; i < len(path)-3; i++ {
		if !isCorner(path[i-1], path[i], path[i+1]) {
			continue
		}
		corners = push(corners, path[i], 3)
		if len(corners) < 3 {
			continue
		}

		fourth, newObj := getFourth(corners[0], corners[1], corners[2])
		if newObj != start &&
			!slices.Contains(path[:i], newObj) &&
			isClear(grid, path[i+1], path[i+2]) &&
			isClear(grid, path[i+2], fourth) {
			result++
		}
	}
	return result
}

func getFourth(first Coord, second Coord, third Coord) (Coord, Coord) {
	var fourth Coord
	var object Coord

	if first.row == second.row {
		fourth = Coord{
			row: third.row,
			col: first.col,
		}
		if fourth.col > third.col {
			object = Coord{
				row: fourth.row,
				col: fourth.col + 1,
			}
		} else {
			object = Coord{
				row: fourth.row,
				col: fourth.col - 1,
			}
		}
	} else {
		fourth = Coord{
			row: third.row,
			col: first.col,
		}
		if fourth.row > third.row {
			object = Coord{
				row: fourth.row + 1,
				col: fourth.col,
			}
		} else {
			object = Coord{
				row: fourth.row - 1,
				col: fourth.col,
			}
		}
	}
	return fourth, object
}

func isCorner(prev Coord, curr Coord, next Coord) bool {
	return prev.row != curr.row && curr.col != next.col ||
		prev.col != curr.col && curr.row != next.row
}

func push(arr []Coord, el Coord, maxLength int) []Coord {
	if len(arr) < maxLength {
		arr = append(arr, el)
	} else {
		arr = append(arr[1:maxLength], el)
	}
	return arr
}

func isClear(grid [][]rune, first Coord, second Coord) bool {
	var path []rune
	if first.row == second.row {
		if first.col <= second.col {
			path = grid[first.row][first.col:second.col]
		} else {
			path = grid[first.row][second.col:first.col]
		}
	} else if first.col == second.col {
		if first.row <= second.row {
			path = verticalSlice(grid, first.col, first.row, second.row)
		} else {
			path = verticalSlice(grid, first.col, second.row, first.row)
		}
	} else {
		return false
	}

	for _, v := range path {
		if v == Object {
			return false
		}
	}
	return true
}
