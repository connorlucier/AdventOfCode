def part1(filename):
	with open(filename, 'r') as file:
		data = [(l.split()[0], int(l.split()[1])) for l in file.read().splitlines()]

	x = 0
	y = 0

	for direction, value in data:
		if direction == 'up':
			y -= value
		elif direction == 'down':
			y += value
		elif direction == 'forward':
			x += value

	return x * y

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)