package main

func part2(l []int, r []int) int {
	result := 0
	for _, lval := range l {
		count := Count(r, lval)
		result += lval * count
	}
	return result
}

func Count(arr []int, target int) int {
	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	return count
}
