import numpy as np

def build_grid(filename):
	origins = []
	sizes = []
	max_col = 0
	max_row = 0

	with open(filename, 'r') as infile:
		for line in infile.readlines():
			claim = line[:-1].split(' ')

			# origin[0] = starting column
			# origin[1] = starting row
			origin = tuple(int(x) for x in claim[2][:-1].split(','))

			# size[0] = number of cols
			# size[1] = number of rows
			size = tuple(int(y) for y in claim[3].split('x'))

			origins.append(origin)
			sizes.append(size)

			max_col = max(max_col, origin[0] + size[0])
			max_row = max(max_row, origin[1] + size[1])

	dim = max(max_col, max_row)
	grid = np.zeros((dim,dim), int)

	for i in range(len(origins)):
		origin = origins[i]
		size = sizes[i]

		for row in range(origin[1], origin[1] + size[1]):
			for col in range(origin[0], origin[0] + size[0]):
				grid[row][col] += 1

	return grid, origins, sizes