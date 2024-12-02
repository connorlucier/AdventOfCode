package main

func part1(values [][]int) int {
	result := 0
	for _, row := range values {
		safe := true
		dir := None
		prev := row[0]
		for _, val := range row[1:] {
			isValid, newDir := validate(prev, val, dir)
			if !isValid {
				safe = false
				break
			}
			dir = newDir
			prev = val
		}
		if safe {
			result++
		}
	}

	return result
}
