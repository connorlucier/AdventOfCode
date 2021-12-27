from point import point
from line import line

def part2(filename):
	with open(filename, 'r') as file:
		raw = file.read().splitlines()

	lines = parse_input(raw)
	grid = get_grid(lines)

	for l in lines:
		for pt in l.get_points():
			grid[pt.y][pt.x] += 1

	flattened = [num for row in grid for num in row if num > 1]
	return len(flattened)

def parse_input(lines):
	result = []

	for l in lines:
		start, end = l.split(' -> ')
		x1, y1 = [int(_) for _ in start.split(',')]
		x2, y2 = [int(_) for _ in end.split(',')]

		pt1 = point(x1, y1)
		pt2 = point(x2, y2)
		result.append(line(pt1, pt2))

	return result

def get_grid(data):
	max_x = max([l.max_x() for l in data])
	max_y = max([l.max_y() for l in data])

	result = []
	for y in range(max_y + 1):
		result.append([])
		for x in range(max_x + 1):
			result[y].append(0)

	return result

if __name__ == '__main__':
	result = part2('input.txt')
	print(result)