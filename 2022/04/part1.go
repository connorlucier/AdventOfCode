package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	total := 0
	for scanner.Scan() {
		var min1, max1, min2, max2 int
		line := scanner.Text()
		fmt.Sscanf(line, "%d-%d,%d-%d", &min1, &max1, &min2, &max2)

		if min1 <= min2 && max1 >= max2 || min2 <= min1 && max2 >= max1 {
			total += 1
		}
	}

	fmt.Println(total)
}
