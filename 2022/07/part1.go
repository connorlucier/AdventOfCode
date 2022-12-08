package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() *node {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var root *node
	var cur *node

	for scanner.Scan() {
		var cmd, name, ftype string

		line := scanner.Text()
		n, _ := fmt.Sscanf(line, "$ %s %s", &cmd, &name)

		if n > 0 {
			if cmd == "ls" {
				continue
			}

			if root == nil {
				root = &node{
					Name:   name,
					IsDir:  true,
					Parent: nil,
				}
				cur = root
			} else if name == ".." {
				cur = cur.Parent
			} else {
				cur = cur.GetChild(name)
			}
		} else {
			fmt.Sscanf(line, "%s %s", &ftype, &name)
			child := &node{
				Name:   name,
				IsDir:  ftype == "dir",
				Parent: cur,
			}

			size, err := strconv.Atoi(ftype)
			if err == nil {
				child.Size = size
			}

			cur.Children = append(cur.Children, child)
		}
	}

	root.CalcSize()
	fmt.Println(root.SumDirsUnder(100000))
	return root
}
