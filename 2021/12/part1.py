from node import node

def part1(filename):
	with open(filename, 'r') as file:
		nodes = parse_input(file.read().splitlines())

	state = { n.name: False for n in nodes }
	start, end = nodes[0:2]
	dfs(start, state)
	return len(nodes)

def dfs(start, state, path=[]):
	state[start.name] = True
	path.append(start)
	for node in start.neighbors:
		print([n.name for n in path])
		if not state.get(node.name):
			dfs(node, state)


def parse_input(lines):
	result_dict = {}
	for line in lines:
		n1, n2 = [n for n in line.split('-')]

		if not result_dict.get(n1):
			result_dict[n1] = node(n1)
		if not result_dict.get(n2):
			result_dict[n2] = node(n2)

		result_dict[n1].add(result_dict[n2])

	return sorted(list(result_dict.values()), key=lambda n: len(n.name), reverse=True)

if __name__ == '__main__':
	result = part1('easy.txt')
	print(result)