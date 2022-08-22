'''
Day 2, part 2:

Takes the same input and finds the claim that does not overlap with any others.

'''

from shared import build_grid

def region_valid(grid, origin, size):
	for row in range(origin[1], origin[1] + size[1]):
		for col in range(origin[0], origin[0] + size[0]):
			if grid[row][col] != 1:
				return False
	return True

def find_claim_id(filename, origin):
	with open(filename, 'r') as infile:
		for line in infile.readlines():
			claim = line[:-1].split(' ')

			if origin == tuple(int(x) for x in claim[2][:-1].split(',')):
				return claim[0]

def part2(filename):
	grid, origins, sizes = build_grid(filename)

	while len(origins) >= 1:
		origin = origins.pop()
		size = sizes.pop()

		if region_valid(grid, origin, size):
			return find_claim_id(filename, origin)

	return None

if __name__ == '__main__':
    claim = part2('input.txt')
    print(claim)