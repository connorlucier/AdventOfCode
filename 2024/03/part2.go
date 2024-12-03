package main

func part2(instructions []Instruction) int {
	result := 0
	isEnabled := true
	for _, instr := range instructions {
		if instr.op == Dont {
			isEnabled = false
		} else if instr.op == Do {
			isEnabled = true
		} else if isEnabled {
			result += instr.Result()
		}
	}
	return result
}
