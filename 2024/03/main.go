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
	re := regexp.MustCompile(`(do|don't|mul)\((([0-9]{1,3}),([0-9]{1,3}))?\)`)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	result := []Instruction{}
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for _, m := range matches {
			op := m[1]
			left, _ := strconv.Atoi(m[3])
			right, _ := strconv.Atoi((m[4]))
			result = append(result, Instruction{
				op,
				left,
				right,
			})
		}
	}
	return result
}

const (
	Mul  = "mul"
	Do   = "do"
	Dont = "don't"
)

type Instruction struct {
	op    string
	left  int
	right int
}

func (i Instruction) Result() int {
	if i.op != Mul {
		return 0
	} else {
		return i.left * i.right
	}
}
