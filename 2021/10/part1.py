def part1(filename):
	with open(filename, 'r') as file:
		lines = file.read().splitlines()

	char_dict = {
		')': ('(', 3),
		']': ('[', 57),
		'}': ('{', 1197),
		'>': ('<', 25137)
	}

	score = 0
	for line in lines:
		stack = []
		for char in line:
			if not char_dict.get(char):
				stack.append(char)
			else:
				expected, penalty = char_dict[char]
				if stack.pop() != expected:
					score += penalty
					break

	return score

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)
