package main

import (
	"fmt"
)

func part1(input []instruction) {
	signalStrength := 0
	x := 1

	for i, instr := range input {
		cycle := i + 1
		if (cycle-20)%40 == 0 {
			signalStrength += x * (cycle)
		}

		if instr.Name == "addx" {
			x += instr.Value
		}
	}

	fmt.Println(signalStrength)
}
