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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		first := line[0:mid]
		second := line[mid:]
		ch := getCommonLetter(first, second)
		sum += getPriority(ch)
	}

	fmt.Println(sum)
}
