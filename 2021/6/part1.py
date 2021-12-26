def part1(filename, days):
	with open(filename, 'r') as file:
		data = [int(val) for val in file.read().split(',')]

	for i in range(days):
		data = update(data)

	return len(data)

def update(data):
	updated = [num - 1 for num in data]

	for i in range(len(data)):
		if data[i] == 0:
			updated[i] = 6
			updated.append(8)

	return updated

if __name__ == '__main__':
	result = part1('input.txt', 80)
	print(result)