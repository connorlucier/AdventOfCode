class node:
	def __init__(self, name):
		self.name = name
		self.neighbors = []

	def __str__(self):
		spacing = ' ' * (len(self.name) + 3)
		return '({}) --> '.format(self.name) + '\n{}--> '.format(spacing).join([v.name for v in self.neighbors])

	def add(self, node):
		if node not in self.neighbors:
			self.neighbors.append(node)
		if self not in node.neighbors:
			node.neighbors.append(self)