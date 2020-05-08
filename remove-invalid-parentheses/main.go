package main

import "fmt"

const LEFT_PAREN = '('
const RIGHT_PAREN = ')'

type Pair struct {
	first  byte
	second byte
}

type Result struct {
	list []string
}

func removeInvalidParentheses(s string) []string {
	pair := &Pair{LEFT_PAREN, RIGHT_PAREN}
	result := &Result{[]string{}}
	remove(s, 0, 0, pair, result)
	return result.list
}

func remove(s string, start int, lastDelete int, pair *Pair, result *Result) {
	cnt := 0
	for i := start; i < len(s); i++ {
		if s[i] == pair.first {
			cnt++
		} else if s[i] == pair.second {
			cnt--
		}
		if cnt < 0 {
			for j := lastDelete; j <= i; j++ {
				if s[j] == pair.second && (j == lastDelete || s[j-1] != pair.second) {
					remove(s[0:j]+s[j+1:len(s)], i, j, pair, result)
				}
			}
			return
		}
	}
	s = reverse(s)
	if pair.first == '(' {
		pair = &Pair{RIGHT_PAREN, LEFT_PAREN}
		remove(s, 0, 0, pair, result)
	} else {
		result.list = append(result.list, s)
	}
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	s := "(())())"
	answers := removeInvalidParentheses(s)
	fmt.Println(answers)
}
