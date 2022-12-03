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

	group := make([]string, 3)
	count := 0
	sum := 0
	for scanner.Scan() {
		group[count] = scanner.Text()
		count = (count + 1) % 3

		if count == 0 {
			ch := getCommonLetter(group...)
			sum += getPriority(ch)
		}
	}

	fmt.Println(sum)
}
