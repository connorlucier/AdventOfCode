package main

func part2(grid [][]rune, aLocs []Coord) int {
	result := 0
	for _, coord := range aLocs {
		x := getX(grid, coord.r, coord.c)
		if x.isXmas() {
			result++
		}
	}
	return result
}

func getX(grid [][]rune, r int, c int) X {
	if r < 1 || r >= len(grid)-1 || c < 1 || c >= len(grid[0])-1 {
		return X{}
	}
	return X{
		upL:   grid[r-1][c-1],
		upR:   grid[r-1][c+1],
		downL: grid[r+1][c-1],
		downR: grid[r+1][c+1],
	}
}

type X struct {
	upL   rune
	upR   rune
	downL rune
	downR rune
}

func (x X) isXmas() bool {
	// up   down  right  left
	// S.S  M.M   M.S    S.M
	// .A.  .A.   .A.    .A.
	// M.M  S.S   M.S    S.M
	return x.upL == 'S' && x.upR == 'S' && x.downL == 'M' && x.downR == 'M' ||
		x.upL == 'M' && x.upR == 'M' && x.downL == 'S' && x.downR == 'S' ||
		x.upL == 'M' && x.upR == 'S' && x.downL == 'M' && x.downR == 'S' ||
		x.upL == 'S' && x.upR == 'M' && x.downL == 'S' && x.downR == 'M'
}
