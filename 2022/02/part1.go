package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	var scoreKey = map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
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
