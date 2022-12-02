package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2() {
	var scoreKey = map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	score := 0
	for scanner.Scan() {
		match := scanner.Text()
		score += scoreKey[match]
	}

	fmt.Println(score)
}
