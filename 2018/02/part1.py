'''
Day 2, part 1:

Takes in an input text file that contains a list of strings, with one string per
line. The result is calculated by counting the number of strings in the input file
containing exactly two or three of the same letter and multiplying the counts
together.
'''
def part1(filename):
	two_count = 0
	three_count = 0

	with open(filename, 'r') as infile:
		for line in infile.readlines():
			has_two = False
			has_three = False

			for c in line:
				if line.count(c) == 2:
					has_two = True
				elif line.count(c) == 3:
					has_three = True

			if has_two:
				two_count += 1
			if has_three:
				three_count += 1

	return two_count * three_count

if __name__ == '__main__':
    count = part1('input.txt')
    print(count)