def part1(filename):
	with open(filename, 'r') as file:
		lines = file.read().splitlines()

	data = parse_input(lines)
	filtered = [d for d in data if d[0] == d[2] or d[1] == d[3]]
	grid = get_grid(filtered)

	for li in filtered:
		start_x = min(li[0], li[2])
		end_x = max(li[0], li[2])
		start_y = min(li[1], li[3])
		end_y = max(li[1], li[3])

		for y in range(start_y, end_y + 1):
			for x in range(start_x, end_x + 1):
				grid[y][x] += 1

	flattened = [num for row in grid for num in row if num > 1]
	return len(flattened)

def parse_input(lines):
	result = []

	for line in lines:
		pt1, pt2 = line.split(' -> ')
		x1, y1 = [int(_) for _ in pt1.split(',')]
		x2, y2 = [int(_) for _ in pt2.split(',')]
		result.append((x1, y1, x2, y2))

	return result

def get_grid(data):
	max_x = max(max([x[0] for x in data]), max([x[2] for x in data]))
	max_y = max(max([y[1] for y in data]), max([y[3] for y in data]))

	result = []
	for y in range(max_y + 1):
		result.append([])
		for x in range(max_x + 1):
			result[y].append(0)

	return result

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)