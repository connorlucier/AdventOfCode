package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	left, right := parseInput()
	p1 := part1(left, right)
	p2 := part2(left, right)
	fmt.Println(p1, p2)
}

func parseInput() ([]int, []int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	left := []int{}
	right := []int{}
	for scanner.Scan() {
		vals := strings.Fields(scanner.Text())
		lval, _ := strconv.Atoi(vals[0])
		rval, _ := strconv.Atoi(vals[1])
		left = append(left, lval)
		right = append(right, rval)
	}

	return left, right
}
