def part2(filename, days):
	with open(filename, 'r') as file:
		data = [int(val) for val in file.read().split(',')]

	init_state = [0] * 9

	for fish in data:
		init_state[fish] += 1

	return calc_growth(init_state, days)

def calc_growth(state, days):
	while days > 0:
		next_state = [0] * 9

		for i in range(9):
			if i == 0:
				next_state[8] = state[i]
				next_state[6] += state[i]
			else:
				next_state[i - 1] += state[i]

		state = next_state
		days -= 1

	return sum(state)

if __name__ == '__main__':
	result = part2('input.txt', 256)
	print(result)