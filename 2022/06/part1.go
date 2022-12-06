package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	scanner.Scan()
	input := scanner.Text()

	for i := range input[:len(input)-4] {
		distinct := true
		for _, c := range input[i : i+4] {
			if strings.Count(input[i:i+4], string(c)) > 1 {
				distinct = false
				break
			}
		}

		if distinct {
			fmt.Println(i + 4)
			break
		}
	}
}
