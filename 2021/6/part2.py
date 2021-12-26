def part2(filename, days):
	with open(filename, 'r') as file:
		data = [int(val) for val in file.read().split(',')]

	init_state = [0] * 9

	for fish in data:
		init_state[fish] += 1

	return calc_growth(init_state, days)

def calc_growth(state, days):
	if days <= 0:
		return sum(state)

	next_state = [0] * 9

	vals = list(range(9))
	vals.reverse()

	for i in vals:
		if i == 0:
			next_state[8] = state[i]
			next_state[6] += state[i]
		else:
			next_state[i - 1] = state[i]

	return calc_growth(next_state, days - 1)

if __name__ == '__main__':
	result = part2('input.txt', 256)
	print(result)