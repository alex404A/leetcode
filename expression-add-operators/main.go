package main

import (
	"fmt"
	"strconv"
)

type Wrapper struct {
	result []string
}

func (wrapper *Wrapper) add(r string) {
	wrapper.result = append(wrapper.result, r)
}

func addOperators(num string, target int) []string {
	if len(num) == 0 {
		return []string{}
	}
	wrapper := &Wrapper{[]string{}}
	find(num, target, 0, "", wrapper)
	return wrapper.result
}

func find(num string, target int, previous int, path string, wrapper *Wrapper) {
	if len(num) == 1 {
		checkAtEnd(num, target, previous, path, wrapper)
		return
	}
	for i := 1; i < len(num); i++ {
		prefix := num[0:i]
		prefixNum, _ := strconv.Atoi(prefix)
		if path == "" {
			find(num[i:len(num)], target-prefixNum, prefixNum, num[0:i], wrapper)
		} else {
			find(num[i:len(num)], target-prefixNum, prefixNum, path+"+"+num[0:i], wrapper)
			find(num[i:len(num)], target+prefixNum, 0-prefixNum, path+"-"+num[0:i], wrapper)
			find(num[i:len(num)], target+previous-previous*prefixNum, previous*prefixNum, path+"*"+num[0:i], wrapper)
		}
		if num[0] == '0' {
			break
		}
	}
	if num[0] != '0' {
		checkAtEnd(num, target, previous, path, wrapper)
	}
}

func checkAtEnd(num string, target int, previous int, path string, wrapper *Wrapper) {
	val, _ := strconv.Atoi(num)
	if path == "" {
		if val == target {
			wrapper.add(num)
		}
		return
	}
	if val == target {
		wrapper.add(path + "+" + num)
	}
	if val == 0-target {
		wrapper.add(path + "-" + num)
	}
	if target+previous == previous*val {
		wrapper.add(path + "*" + num)
	}
}

func main() {
	fmt.Println(addOperators("123", 6))
	fmt.Println(addOperators("232", 8))
	fmt.Println(addOperators("105", 5))
	fmt.Println(addOperators("00", 0))
	fmt.Println(addOperators("3456237490", 9191))
	fmt.Println(addOperators("1111", 0))
}
