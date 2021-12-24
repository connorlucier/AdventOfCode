def part1(filename):
	file = open(filename, 'r')
	lines = file.readlines()

	x = 0
	y = 0

	for l in lines:
		(direction, _) = l.split(' ')
		value = int(_)

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