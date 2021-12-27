def part2(filename):
	with open(filename, 'r') as file:
		data = [(l.split()[0], int(l.split()[1])) for l in file.read().splitlines()]

	a = 0
	x = 0
	y = 0

	for direction, value in data:
		if direction == 'up':
			a -= value
		elif direction == 'down':
			a += value
		elif direction == 'forward':
			x += value
			y += a * value

	return x * y

if __name__ == '__main__':
	result = part2('input.txt')
	print(result)