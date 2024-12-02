package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	None       = "none"
	Ascending  = "asc"
	Descending = "desc"
)

func main() {
	values := parseInput()
	fmt.Println(part1(values), part2(values))
}

func parseInput() [][]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	result := [][]int{}
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		row := []int{}
		for _, v := range strs {
			num, _ := strconv.Atoi(v)
			row = append(row, num)
		}
		result = append(result, row)
	}
	return result
}

func isRowSafe(row []int) bool {
	isSafe := true
	dir := None
	prev := row[0]
	for _, val := range row[1:] {
		isSafe, dir = isValueSafe(prev, val, dir)
		if !isSafe {
			break
		}
		prev = val
	}
	return isSafe
}

func isValueSafe(prev int, val int, dir string) (bool, string) {
	if prev == val {
		return false, dir
	}
	diff := int(math.Abs(float64(val - prev)))
	if diff > 3 || diff < 1 {
		return false, dir
	} else if (dir == Ascending && val < prev) || (dir == Descending && val > prev) {
		return false, dir
	}
	newDir := dir
	if dir == None {
		if val > prev {
			newDir = Ascending
		} else {
			newDir = Descending
		}
	}
	return true, newDir
}
