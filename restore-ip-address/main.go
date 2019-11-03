package main

import (
	"fmt"
	"strconv"
)

type container struct {
	values []string
}

func restoreIpAddresses(s string) []string {
	empty := make([]string, 0)
	if len(s) <= 3 {
		return empty
	}
	if len(s) > 12 {
		return empty
	}
	result := container{make([]string, 0)}
	find(s, 4, "", &result)
	return result.values
}

func find(s string, num int, prefix string, values *container) {
	if num == 1 {
		check(s, prefix, values)
	}
	if len(s) > 1 {
		prefix1 := supplement(prefix, s[:1])
		find(s[1:], num-1, prefix1, values)
	}
	if len(s) > 2 && s[0] != '0' {
		prefix2 := supplement(prefix, s[:2])
		find(s[2:], num-1, prefix2, values)
	}
	if len(s) > 3 && s[0] != '0' {
		value, _ := strconv.Atoi(s[:3])
		if value <= 255 {
			prefix3 := supplement(prefix, s[:3])
			find(s[3:], num-1, prefix3, values)
		}
	}
}

func supplement(prefix string, next string) string {
	prefix = prefix + next + "."
	return prefix
}

func check(s string, prefix string, values *container) {
	if len(s) > 3 || len(s) == 0 {
		return
	} else if len(s) <= 2 {
		if len(s) == 2 && s[0] == '0' {
			return
		}
		prefix += s
		values.values = append(values.values, prefix)
	} else if len(s) == 3 {
		if s[0] == '0' {
			return
		}
		num, _ := strconv.Atoi(s)
		if num <= 255 {
			prefix += s
			values.values = append(values.values, prefix)
		} else {
			return
		}
	}
}

func main() {
	str := "00000"
	result := restoreIpAddresses(str)
	fmt.Println(result)
}
