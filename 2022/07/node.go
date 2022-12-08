package main

type node struct {
	Name     string
	Size     int
	IsDir    bool
	Parent   *node
	Children []*node
}

func (n node) GetChild(name string) *node {
	for _, c := range n.Children {
		if c.Name == name {
			return c
		}
	}

	return nil
}

func (n *node) CalcSize() int {
	if !n.IsDir {
		return n.Size
	}

	size := 0
	for _, c := range n.Children {
		size += c.CalcSize()
	}

	n.Size = size
	return size
}

func (n node) SumDirsUnder(maxSize int) int {
	total := 0
	for _, c := range n.Children {
		if !c.IsDir {
			continue
		}
		if c.Size <= maxSize {
			total += c.Size
		}

		total += c.SumDirsUnder(maxSize)
	}

	return total
}

func (n *node) FindWithSize(minSize int) *node {
	if !n.IsDir || n.Size < minSize {
		return nil
	}

	minDir := n
	for _, c := range n.Children {
		if !c.IsDir {
			continue
		}

		if c.Size >= minSize {
			childMin := c.FindWithSize(minSize)
			if childMin != nil && childMin.Size < minDir.Size {
				minDir = childMin
			} else if c.Size < minDir.Size {
				minDir = c
			}
		}
	}

	return minDir
}
