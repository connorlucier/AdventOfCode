def part1(filename):
	with open(filename, 'r') as file:
		data = [int(val) for val in file.read().split(',')]

	max_pos = max(data)
	min_pos = min(data)

	fuel_to_pos = [0] * (max_pos - min_pos)
	for i in range(len(fuel_to_pos)):
		fuel_to_pos[i] = sum([abs(val - i) for val in data])

	return min(fuel_to_pos)

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)