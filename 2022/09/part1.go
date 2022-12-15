package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	r := NewRope(2)
	var visited []point

	for scanner.Scan() {
		line := scanner.Text()
		var dir string
		var num int

		fmt.Sscanf(line, "%s %d", &dir, &num)

		for i := 0; i < num; i++ {
			r.Move(dir)
			if !Contains(visited, *r.Tail) {
				visited = append(visited, *r.Tail)
			}
		}
	}

	fmt.Println(len(visited))
}
