package main

func part2(instructions []Instruction) int {
	result := 0
	for _, instr := range instructions {
		if instr.isEnabled {
			result += instr.Result()
		}
	}
	return result
}
