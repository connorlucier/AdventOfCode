from color import pretty_print

def part2(filename):
	with open(filename, 'r') as file:
		state = [[int(i) for i in l] for l in file.read().splitlines()]

	step_count = 0
	flash_count = 0
	while flash_count < len(state) ** 2:
		flash_pts = []

		for i in range(len(state)):
			for j in range(len(state[i])):
				state[i][j] = increment(state[i][j])

				if state[i][j] == 0:
					flash_pts.append((i, j))

		flash_neighbors(flash_pts, state)
		# pretty_print(state)

		flash_count = len([num for row in state for num in row if num == 0])
		step_count += 1

	return step_count

def increment(val):
	return (val + 1) % 10

def flash_neighbors(pts, state):
	if not pts:
		return

	new_flash_pts = []
	for pt in pts:
		for (i, j) in get_neighbors(state, pt):
			if state[i][j] != 0:
				state[i][j] = increment(state[i][j])
				if state[i][j] == 0:
					new_flash_pts.append((i,j))

	flash_neighbors(new_flash_pts, state)

def get_neighbors(data, pt):
	i, j = pt
	result = []

	if i > 0:
		result.append((i-1, j))
		if j > 0:
			result.append((i-1, j-1))
		if j < len(data[i]) - 1:
			result.append((i-1, j+1))
	if i < len(data) - 1:
		result.append((i+1, j))
		if j > 0:
			result.append((i+1, j-1))
		if j < len(data[i]) - 1:
			result.append((i+1, j+1))
	if j > 0:
		result.append((i, j-1))
	if j < len(data[i]) - 1:
		result.append((i, j+1))

	return result

if __name__ == '__main__':
	result = part2('input.txt')
	print(result)
