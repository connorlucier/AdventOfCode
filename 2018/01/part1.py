'''
Day 1, part 1:

Takes in an input text file that contains a list of numbers, with one number each
line. Beginning at zero, the result is calculated by adding the value of each line
to the running total. The format of each line is a sign character followed by the
number.
'''
def part1(filename):
	freq = 0

	with open(filename, 'r') as infile:
		for line in infile.readlines():
			if line[0] == '+':
				freq += int(line[1:])
			else:
				freq += int(line)
	return freq

if __name__ == '__main__':
    freq = part1('input.txt')
    print(freq)