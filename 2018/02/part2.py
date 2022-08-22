'''
Day 2, part 2:

Takes the same input text file and finds the first pair of strings that differ by
only one character in the same position of both strings.

'''
def part2(filename):
	with open(filename, 'r') as infile:
		lines = infile.readlines()

		for i in range(len(lines)):
			a = lines[i][:-1]

			for j in range(i+1, len(lines)):
				b = lines[j][:-1]

				for k in range(len(b)):
					temp_a = a[:k] + a[k+1:]
					temp_b = b[:k] + b[k+1:]

					if temp_a == temp_b:
						return temp_a
	return ''

if __name__ == '__main__':
	output = part2('input.txt')
	print(output)
