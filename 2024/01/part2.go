package main

func part2(l []int, r []int) int {
	result := 0
	for _, lval := range l {
		count := 0
		for _, rval := range r {
			if lval == rval {
				count++
			}
		}
		result += lval * count
	}

	return result
}
