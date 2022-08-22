'''
Day 1, part 2:

Finds the first repeated frequency in the visited frequencies.
'''
def part2(filename):
	freq = 0
	freqs = [0]

	while True:
		with open(filename, 'r') as infile:
			for line in infile.readlines():
				if line[0] == '+':
					freq += int(line[1:])
				else:
					freq += int(line)

				if freq in freqs:
					return freq

				freqs.append(freq)

if __name__ == '__main__':
    freq = part2('input.txt')
    print(freq)