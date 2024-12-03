package main

func part1(instructions []Instruction) int {
	result := 0
	for _, instr := range instructions {
		result += instr.Result()
	}
	return result
}
