def part2(filename):
	with open(filename, 'r') as file:
		data = [int(l) for l in file.read().splitlines()]

	count = 0

	for i in range(1, len(data) - 2):
		prev = sum(data[i-1:i+2])
		curr = sum(data[i:i+3])

		if curr > prev:
			count += 1

	return count

if __name__ == '__main__':
	count = part2('input.txt')
	print(count)