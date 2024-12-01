package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

func part2() {
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

	result := 0
	for _, lval := range left {
		count := 0
		for _, rval := range right {
			if lval == rval {
				count++
			}
		}

		result += lval * count
	}

	fmt.Println(result)
}
