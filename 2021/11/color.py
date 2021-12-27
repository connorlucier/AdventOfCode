class color:
	purple = '\033[95m'
	dark_cyan = '\033[36m'
	cyan = '\033[96m'
	green = '\033[92m'
	yellow = '\033[93m'
	red = '\033[91m'
	end = '\033[0m'

	key = {
		0: purple,
		1: dark_cyan,
		2: dark_cyan,
		3: cyan,
		4: cyan,
		5: green,
		6: green,
		7: yellow,
		8: yellow,
		9: red
	}

def pretty_print(data):
	print('\n'.join([''.join([color.key[c] + str(c) + color.end for c in r]) for r in data]))
	print('\n')
