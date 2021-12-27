def part2(filename):
	with open(filename, 'r') as file:
		lines = file.read().splitlines()

	char_dict = {
		'(': (')', 1),
		'[': (']', 2),
		'{': ('}', 3),
		'<': ('>', 4)
	}

	inverted_dict = { v[0]: k for k, v in char_dict.items() }
	filtered = list(filter(lambda l: not is_corrupted(l, inverted_dict), lines))

	scores = []
	for line in filtered:
		stack = []
		for char in line:
			if char_dict.get(char):
				stack.append(char)
			else:
				stack.pop()

		scores.append(calc_score(reversed(stack), char_dict))

	scores.sort()
	mid = len(scores) // 2
	return scores[mid]

def is_corrupted(line, dct):
	stack = []
	for char in line:
		if not dct.get(char):
			stack.append(char)
		else:
			expected = dct[char]
			if stack.pop() != expected:
				return True

	return False

def calc_score(stack, dct):
	score = 0
	for c in stack:
		score *= 5
		score += dct[c][1]

	return score

if __name__ == '__main__':
	result = part2('input.txt')
	print(result)
