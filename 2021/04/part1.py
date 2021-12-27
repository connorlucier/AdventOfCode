def part1(filename):
	with open(filename, 'r') as file:
		lines = file.read().splitlines()

	sequence, boards = parse_input(lines)
	board, seq = get_first_bingo(boards, sequence)

	flattened = [it for row in board for it in row]
	filtered = [num for num in flattened if num not in seq]

	return sum(filtered) * seq[-1]

def parse_input(data):
	sequence = [int(i) for i in data[0].split(',')]
	boards = []
	b = []

	for row in data[1:]:
		if not row:
			if b:
				boards.append(b)
				b = []
		else:
			b.append([int(i) for i in row.split()])

	return sequence, boards

def check_board(brd, seq):
	for row in brd:
		bingo = True
		for num in row:
			if num not in seq:
				bingo = False
				break

		if bingo:
			return True

	for col in range(len(brd[0])):
		bingo = True
		for row in brd:
			if row[col] not in seq:
				bingo = False
				break
		if bingo:
			return True

	return False

def get_first_bingo(boards, sequence):
	for i in range(5, len(sequence)):
		for brd in boards:
			seq = sequence[0:i]
			bingo = check_board(brd, seq)

			if bingo:
				return brd, seq
	return [], []

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)