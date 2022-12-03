package main

import (
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func getCommonLetter(strs ...string) byte {
	if len(strs) <= 1 {
		panic("must compare two or more strings")
	}

	for _, v := range strs[0] {
		valid := false
		for _, str := range strs[1:] {
			valid = strings.ContainsRune(str, v)
			if !valid {
				break
			}
		}

		if valid {
			return byte(v)
		}
	}

	fmt.Println(strs)
	panic("strings have no common letter")
}

func getPriority(c byte) int {
	if c <= 90 && c >= 65 {
		return int(c) - 38
	} else if c <= 122 && c >= 97 {
		return int(c) - 96
	}

	panic("invalid byte")
}
