def part2(filename):
	with open(filename, 'r') as file:
		data = [[int(c) for c in row] for row in file.read().splitlines()]

	low_points = []
	for i in range(len(data)):
		for j in range(len(data[i])):
			if is_low_point(data, i, j):
				low_points.append((i, j))

	basin_sizes = [calc_basin(data, pt, []) for pt in low_points]
	basin_sizes.sort()
	top_three = basin_sizes[-3:]

	return top_three[0] * top_three[1] * top_three[2]

def calc_basin(data, pt, visited):
	i, j = pt

	if out_of_bounds(data, pt) or data[i][j] >= 9:
		return 0

	visited.append(pt)
	neighbors = filter(lambda s: s not in visited, [(i+1, j), (i-1, j), (i, j+1), (i, j-1)])

	return 1 + sum([calc_basin(data, n, visited) for n in neighbors])

def is_low_point(data, i, j):
	val = data[i][j]
	result = True

	if i > 0:
		result &= val < data[i-1][j]
	if i < len(data) - 1:
		result &= val < data[i+1][j]
	if j > 0:
		result &= val < data[i][j-1]
	if j < len(data[i]) - 1:
		result &= val < data[i][j+1]

	return result

def out_of_bounds(data, pt):
	i, j = pt
	return i < 0 or j < 0 or i >= len(data) or j >= len(data[i])

if __name__ == '__main__':
	result = part2('input.txt')
	print(result)
