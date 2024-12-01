package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"os"
)

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	left := make([]int, 1)
	right := make([]int, 1)
	for scanner.Scan() {
		vals := strings.Fields(scanner.Text())
		lval, _ := strconv.Atoi(vals[0])
		rval, _ := strconv.Atoi(vals[1])
		left = append(left, lval)
		right = append(right, rval)
	}

	slices.Sort(left)
	slices.Sort(right)

	result := 0
	for i, _ := range left {
		diff := float64(left[i] - right[i])
		result += int(math.Abs(diff))
	}

	fmt.Println(result)
}
