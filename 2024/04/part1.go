package main

func part1(grid [][]rune, xLocs []Coord) int {
	result := 0
	target := "XMAS"
	for _, coord := range xLocs {
		strs := getStrings(&grid, coord.r, coord.c, len(target))
		for _, s := range strs {
			if s == target {
				result++
			}
		}
	}
	return result
}

func getStrings(grid *[][]rune, r int, c int, length int) []string {
	height := len(*grid)
	width := len((*grid)[0])
	result := []string{}

	// up
	if r >= length-1 {
		result = append(result, getString(grid, r, r-length, c, c))
		// up + right
		if c <= width-length {
			result = append(result, getString(grid, r, r-length, c, c+length))
		}
		// up + left
		if c >= length-1 {
			result = append(result, getString(grid, r, r-length, c, c-length))
		}
	}
	// down
	if r <= height-length {
		result = append(result, getString(grid, r, r+length, c, c))
		//down + right
		if c <= width-length {
			result = append(result, getString(grid, r, r+length, c, c+length))
		}
		// down + left
		if c >= length-1 {
			result = append(result, getString(grid, r, r+length, c, c-length))
		}
	}
	// right
	if c <= width-length {
		result = append(result, getString(grid, r, r, c, c+length))
	}
	// left
	if c >= length-1 {
		result = append(result, getString(grid, r, r, c, c-length))
	}

	return result
}

func getString(grid *[][]rune, startR int, endR int, startC int, endC int) string {
	rStep := 1
	cStep := 1

	if startR > endR {
		rStep = -1
	} else if startR == endR {
		rStep = 0
	}
	if startC > endC {
		cStep = -1
	} else if startC == endC {
		cStep = 0
	}

	r := startR
	c := startC
	result := []rune{}
	for !(r == endR && c == endC) {
		result = append(result, (*grid)[r][c])
		r += rStep
		c += cStep
	}

	return string(result)
}
