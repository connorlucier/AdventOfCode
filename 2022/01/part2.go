package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func part2() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	maxes := make([]int, 3)
	current := 0

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			min, i := minimum(maxes)
			if current > min {
				maxes[i] = current
			}

			current = 0
			continue
		}

		next, _ := strconv.Atoi(text)
		current += next
	}

	fmt.Println(sum(maxes))
}

func minimum(arr []int) (int, int) {
	if len(arr) == 0 {
		panic(errors.New("array length must be greater than zero"))
	}

	min := arr[0]
	index := 0

	for i, val := range arr {
		if val < min {
			min = val
			index = i
		}
	}

	return min, index
}

func sum(vals []int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}

	return sum
}
