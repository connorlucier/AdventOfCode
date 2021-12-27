def part1(filename):
	with open(filename, 'r') as file:
		data = parse_input(file.readlines())

	length_map = {
		2: 1,
		3: 7,
		4: 4,
		7: 8
	}

	total = 0
	for row in data:
		total += sum([1 if length_map.get(val) else 0 for val in row])

	return total

def parse_input(lines):
	result = []
	for l in lines:
		output = l.split('|')[1].strip()
		result.append([len(val) for val in output.split()])

	return result

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)
