def part1(filename):
	file = open(filename, 'r')
	lines = file.read().splitlines()
	max_len = max([len(l) for l in lines])
	bits = [0] * max_len

	for j in range(max_len):
		zeros = 0
		ones = 0

		for i in range(len(lines)):
			if lines[i][j] == '1':
				ones += 1
			else:
				zeros += 1

		bits[j] = '1' if ones > zeros else '0'

	binary = ''.join(bits)
	gamma = bin_to_dec(binary)
	epsilon = bin_to_dec(flip_bits(binary))

	return gamma * epsilon

def bin_to_dec(val):
	nums = [int(c) for c in val]
	result = 0
	length = len(nums)

	for i in range(length):
		result += nums[i] * 2 ** (length - 1 - i)

	return result

def flip_bits(val):
	return ''.join(['1' if c == '0' else '0' for c in val])

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)