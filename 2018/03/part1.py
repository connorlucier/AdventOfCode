'''
Day 3, part 1:

Takes in an input text file containing a list of "claims" to specific areas of
a grid. Calculates the number of grid spaces with two or more people claiming the
space.
'''
from shared import build_grid

def part1(filename):
	grid, _, _ = build_grid(filename)

	result = 0
	for row in grid:
		result += (len(row) - list(row).count(0) - list(row).count(1))

	return result

if __name__ == '__main__':
	size = part1('input.txt')
	print(size)