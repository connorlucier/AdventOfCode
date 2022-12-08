package main

import "fmt"

func part2(root *node) {
	spaceRequired := 30000000 - (70000000 - root.Size)
	minDirToDelete := root.FindWithSize(spaceRequired)

	if minDirToDelete == nil {
		panic("no dir big enough")
	}

	fmt.Println(minDirToDelete.Size)
}
