def part2(filename):
	file = open(filename, 'r')
	lines = file.readlines()

	a = 0
	x = 0
	y = 0

	for l in lines:
		(direction, _) = l.split(' ')
		value = int(_)

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