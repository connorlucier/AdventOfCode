from point import point

class line:
	def __init__(self, pt1, pt2):
		self.x1 = pt1.x
		self.y1 = pt1.y
		self.x2 = pt2.x
		self.y2 = pt2.y

	def __str__(self):
		return str(self.start()) + ' --> ' + str(self.end())

	def start(self):
		return point(self.x1, self.y1)

	def end(self):
		return point(self.x2, self.y2)

	def min_x(self):
		return self.x1 if self.x1 <= self.x2 else self.x2

	def max_x(self):
		return self.x1 if self.x1 >= self.x2 else self.x2

	def min_y(self):
		return self.y1 if self.y1 <= self.y2 else self.y2

	def max_y(self):
		return self.y1 if self.y1 >= self.y2 else self.y2

	def is_vertical(self):
		return self.x1 == self.x2

	def is_horizontal(self):
		return self.y1 == self.y2

	def is_diagonal(self):
		return not self.is_horizontal() and not self.is_vertical()

	def slope(self):
		if self.is_vertical():
			return None
		if self.is_horizontal():
			return 0

		return (self.y2 - self.y1) / (self.x2 - self.x1)

	def get_points(self):
		if self.is_horizontal():
			return [point(x, self.y1) for x in range(self.min_x(), self.max_x() + 1)]
		elif self.is_vertical():
			return [point(self.x1, y) for y in range(self.min_y(), self.max_y() + 1)]
		else:
			return [point(self.min_x() + i if self.slope() > 0 else self.max_x() - i, self.min_y() + i) for i in range(self.max_y() - self.min_y() + 1)]
