package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	total := 0
	for scanner.Scan() {
		var min1, max1, min2, max2 int
		line := scanner.Text()
		fmt.Sscanf(line, "%d-%d,%d-%d", &min1, &max1, &min2, &max2)

		// "overlaps" --> ((min2 between min1 and max1) OR (max2 between min1 and max1)) OR vice versa
		if (min1 <= min2 && min2 <= max1 || min1 <= max2 && max2 <= max1) || (min2 <= min1 && min1 <= max2 || min2 <= max1 && max1 <= max2) {
			total += 1
		}
	}

	fmt.Println(total)
}
