def part2(filename):
	with open(filename, 'r') as file:
		lines = file.read().splitlines()

	ox = bin_to_dec(filter_data(lines, 0, True))
	co = bin_to_dec(filter_data(lines, 0, False))

	return ox * co

def filter_data(data, pos, takeMax):
	if len(data) <= 0:
		return None
	elif len(data) == 1:
		return data[0]

	zeros = 0
	ones = 0

	for i in range(len(data)):
		if data[i][pos] == '1':
			ones += 1
		else:
			zeros += 1

	bit = '1' if (takeMax and ones >= zeros) or (not takeMax and ones < zeros) else '0'
	return filter_data([d for d in data if d[pos] == bit], pos + 1, takeMax)

def bin_to_dec(val):
	nums = [int(c) for c in val]
	result = 0
	length = len(nums)

	for i in range(length):
		result += nums[i] * 2 ** (length - 1 - i)

	return result

if __name__ == '__main__':
	result = part2('input.txt')
	print(result)
