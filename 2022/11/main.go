package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func parseInput(filename string) ([]*monkey, int) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	defer file.Close()

	monkeyRegex, _ := regexp.Compile(`Monkey ([0-9]+):`)
	itemsRegex, _ := regexp.Compile(`\s+Starting items: (([0-9]+,?\s?)+)`)
	opRegex, _ := regexp.Compile(`\s+Operation: new = old (\+|\*) ([0-9]+|old)`)
	testRegex, _ := regexp.Compile(`\s+Test: divisible by ([0-9]+)`)
	targetRegex, _ := regexp.Compile(`\s+If (true|false): throw to monkey ([0-9]+)`)

	var firstTargets []int
	var secondTargets []int
	var monkeys []*monkey
	var next *monkey
	for scanner.Scan() {
		line := scanner.Text()

		var matches []string
		if monkeyRegex.MatchString(line) {
			matches = monkeyRegex.FindStringSubmatch(line)
			next = &monkey{
				ID: matches[1],
			}
		} else if itemsRegex.MatchString(line) {
			matches = itemsRegex.FindStringSubmatch(line)
			for _, str := range strings.Split(matches[1], ", ") {
				item, _ := strconv.Atoi(str)
				next.Items = append(next.Items, item)
			}
		} else if opRegex.MatchString(line) {
			matches = opRegex.FindStringSubmatch(line)

			if matches[2] == "old" {
				next.Operation = "**"
				next.Operand = 2
			} else {
				operand, _ := strconv.Atoi(matches[2])
				next.Operation = matches[1]
				next.Operand = operand
			}
		} else if testRegex.MatchString(line) {
			matches = testRegex.FindStringSubmatch(line)
			divisor, _ := strconv.Atoi(matches[1])
			next.Divisor = divisor
		} else if targetRegex.MatchString(line) {
			matches = targetRegex.FindStringSubmatch(line)
			first, _ := strconv.ParseBool(matches[1])
			target, _ := strconv.Atoi(matches[2])

			if first {
				firstTargets = append(firstTargets, target)
			} else {
				secondTargets = append(secondTargets, target)
				monkeys = append(monkeys, next)
			}
		}
	}

	for i, m := range monkeys {
		m.FirstTarget = monkeys[firstTargets[i]]
		m.SecondTarget = monkeys[secondTargets[i]]
	}

	nums := make([]int, len(monkeys))
	for i, m := range monkeys {
		nums[i] = m.Divisor
	}

	return monkeys, leastCommonMultiple(nums)
}

func monkeyBusiness(monkeys []*monkey) int {
	max := 0
	second := 0
	for _, m := range monkeys {
		if m.Inspections > max {
			second = max
			max = m.Inspections
		} else if m.Inspections > second {
			second = m.Inspections
		}
	}

	return max * second
}

func leastCommonMultiple(nums []int) int {
	factors := map[int]int{}

	for _, n := range nums {
		primes := map[int]int{}
		for n%2 == 0 {
			if primes[2] == 0 {
				primes[2] = 1
			}
			primes[2] *= 2
			n /= 2
		}

		for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
			for n%i == 0 {
				if primes[i] == 0 {
					primes[i] = 1
				}
				primes[i] *= i
				n /= i
			}
		}

		if n > 2 {
			primes[n] = n
		}

		for k, v := range primes {
			if v > factors[k] {
				factors[k] = v
			}
		}
	}

	result := 1
	for _, v := range factors {
		result *= v
	}

	return result
}
