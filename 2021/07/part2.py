def part2(filename):
	with open(filename, 'r') as file:
		data = [int(val) for val in file.read().split(',')]

	max_pos = max(data)
	min_pos = min(data)

	fuel_to_pos = [0] * (max_pos - min_pos)
	for i in range(len(fuel_to_pos)):
		for val in data:
			steps_to_pos = abs(val - i)
			fuel_to_pos[i] += steps_to_pos * (steps_to_pos + 1) // 2

	return min(fuel_to_pos)

if __name__ == '__main__':
	result = part2('input.txt')
	print(result)
