package main

import "fmt"

func part2() {
	monkeys, lcm := parseInput("input.txt")
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.Act(lcm, false)
		}
	}

	fmt.Println(monkeyBusiness(monkeys))
}
