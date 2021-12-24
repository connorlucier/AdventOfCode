def part1(filename):
	file = open(filename, 'r')
	lines = file.read().splitlines()
	sequence, boards = parse_input(lines)

	for i in range(5, len(sequence)):
		for b in boards:
			check_board(b, sequence[0:i])

def parse_input(data):
	sequence = data[0].split(',')
	boards = []
	b = []

	for row in data[1:]:
		if not row:
			if b:
				boards.append(b)
				b = []
			continue
		b.append([int(i) for i in row.split()])

	return sequence, boards

def check_board(board, seq):
	for row in board:
		bingo = True
		for num in row:
			if num not in seq:
				bingo = False
				break
		if bingo:
			return board

	for col in range(len(board[0])):
		bingo = True
		for row in board:
			if row[col] not in seq:
				bingo = False
				break
		if bingo:
			return board

	return None

if __name__ == '__main__':
	result = part1('input.txt')
	print(result)