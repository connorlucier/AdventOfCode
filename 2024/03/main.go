package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	instructions := parseInput()
	fmt.Println(part1(instructions), part2(instructions))
}

func parseInput() []Instruction {
	instrRe := regexp.MustCompile(`^mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	onRe := regexp.MustCompile(`^do\(\)`)
	offRe := regexp.MustCompile(`^don't\(\)`)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	result := []Instruction{}
	isEnabled := true
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); {
			if onRe.MatchString(line[i:]) {
				isEnabled = true
				i += 4
			} else if offRe.MatchString(line[i:]) {
				isEnabled = false
				i += 7
			} else if instrRe.MatchString(line[i:]) {
				matches := instrRe.FindStringSubmatch(line[i:])
				left, _ := strconv.Atoi(matches[1])
				right, _ := strconv.Atoi(matches[2])
				result = append(result, Instruction{
					left,
					right,
					isEnabled,
				})
				i += len(matches[0])
			} else {
				i++
			}
		}
	}
	return result
}

type Instruction struct {
	left      int
	right     int
	isEnabled bool
}

func (i Instruction) Result() int {
	return i.left * i.right
}
