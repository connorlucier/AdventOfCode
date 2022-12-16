package main

import "fmt"

func part1() {
	monkeys, lcm := parseInput("input.txt")
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.Act(lcm, true)
		}
	}

	fmt.Println(monkeyBusiness(monkeys))
}
