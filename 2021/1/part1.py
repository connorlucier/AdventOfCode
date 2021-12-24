def part1(filename):
	with open(filename, 'r') as file:
		data = [int(l) for l in file.read().splitlines()]

	count = 0

	for i in range(1, len(data)):
		if data[i] > data[i-1]:
			count += 1

	return count

if __name__ == '__main__':
	count = part1('input.txt')
	print(count)