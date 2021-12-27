def part1(filename):
	with open(filename, 'r') as file:
		data = [[int(c) for c in row] for row in file.read().splitlines()]

	total = 0
	for i in range(len(data)):
		for j in range(len(data[i])):
			if is_low_point(data, i, j):
				total += data[i][j] + 1

	return total

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

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)
