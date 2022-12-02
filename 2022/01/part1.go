package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	max := 0
	current := 0

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			if current > max {
				max = current
			}

			current = 0
			continue
		}

		next, _ := strconv.Atoi(text)
		current += next
	}

	fmt.Println(max)
}
