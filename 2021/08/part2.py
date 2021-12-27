def part2(filename):
	with open(filename, 'r') as file:
		data = parse_input(file.readlines())

	total = 0
	for inpt, outpt in data:
		one, seven, four = inpt[0:3]
		eight = inpt[-1]

		codex = {}

		len_five = [i for i in inpt if len(i) == 5]
		len_six = [i for i in inpt if len(i) == 6]

		nine = next(filter(lambda s: s.issuperset(four), len_six))
		len_six.remove(nine)

		three = next(filter(lambda s: s.issuperset(one), len_five))
		len_five.remove(three)

		five = next(filter(lambda s: nine.issuperset(s), len_five))
		len_five.remove(five)

		two = len_five[0]

		six = next(filter(lambda s: s.issuperset(five), len_six))
		len_six.remove(six)

		zero = len_six[0]

		codex['a'] = to_str(seven - one)
		codex['b'] = to_str(eight - (one | two))
		codex['c'] = to_str(eight - six)
		codex['d'] = to_str(eight - zero)
		codex['e'] = to_str(eight - nine)
		codex['f'] = to_str(one - two)
		codex['g'] = to_str(nine - (four | seven))

		decoded = [decode(val, codex) for val in (zero, one, two, three, four, five, six, seven, eight, nine)]
		total += int(''.join([str(decoded.index(decode(val, codex))) for val in outpt]))

	return total

def parse_input(lines):
	result = []
	for l in lines:
		raw_in, raw_out = l.split('|')

		inpt = [set(x) for x in raw_in.strip().split()]
		outpt = [x for x in raw_out.strip().split()]

		inpt.sort(key=len)
		result.append((inpt, outpt))

	return result

def to_str(s):
	return ''.join(s)

def decode(val, codex):
	return set([codex[c] for c in val])


if __name__ == '__main__':
	result = part2('input.txt')
	print(result)
