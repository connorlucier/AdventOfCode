package main

import "fmt"

func part2(input []instruction) {
	screen := make([][]string, 6)
	for i := range screen {
		screen[i] = make([]string, 40)
	}

	x := 1
	for i, instr := range input {
		row := i / 40
		col := i % 40

		if col <= x+1 && col >= x-1 {
			screen[row][col] = "#"
		} else {
			screen[row][col] = "."
		}

		if instr.Name == "addx" {
			x += instr.Value
		}
	}

	render(screen)
}

func render(screen [][]string) {
	for _, row := range screen {
		for _, pixel := range row {
			fmt.Print(pixel)
		}
		fmt.Println()
	}
}
