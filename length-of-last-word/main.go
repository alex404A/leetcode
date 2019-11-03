package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	count := 0
	last := len(s) - 1

	for i := last; i >= 0; i-- {
		c := s[i]
		if c != ' ' {
			break
		} else {
			last--
		}
	}

	for i := last; i >= 0; i-- {
		c := s[i]
		if c != ' ' {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(lengthOfLastWord(""))
}
