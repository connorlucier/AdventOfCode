def part1(filename):
	with open(filename, 'r') as file:
		data = [int(l) for l in file.read().splitlines()]
	return sum([1 if data[i] > data[i-1] else 0 for i in range(1, len(data))])

if __name__ == '__main__':
	count = part1('input.txt')
	print(count)