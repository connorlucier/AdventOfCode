package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part2() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	scanner.Scan()
	input := scanner.Text()

	for i := range input[:len(input)-14] {
		distinct := true
		for _, c := range input[i : i+14] {
			if strings.Count(input[i:i+14], string(c)) > 1 {
				distinct = false
				break
			}
		}

		if distinct {
			fmt.Println(i + 14)
			break
		}
	}
}
