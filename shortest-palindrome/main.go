package main

import "fmt"

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s)%2 == 0 {
		return shortestEven(s)
	} else {
		return shortestOdd(s)
	}
}

func shortestOdd(s string) string {
	pivot := len(s) / 2
	for pivot >= 0 {
		result := checkSingle(pivot, s)
		if result {
			return supplementSingle(pivot, s)
		}
		pivot--
		result = checkDouble(pivot, s)
		if result {
			return supplementDouble(pivot, s)
		}
	}
	return supplementSingle(0, s)
}

func shortestEven(s string) string {
	pivot := len(s)/2 - 1
	for pivot >= 0 {
		result := checkDouble(pivot, s)
		if result {
			return supplementDouble(pivot, s)
		}
		result = checkSingle(pivot, s)
		if result {
			return supplementSingle(pivot, s)
		}
		pivot--
	}
	return supplementSingle(0, s)
}

func supplementSingle(pivot int, s string) string {
	start := 2*pivot + 1
	bytes := make([]byte, len(s)-start)
	for i := start; i < len(s); i++ {
		bytes[len(s)-1-i] = s[i]
	}
	return string(bytes) + s
}

func supplementDouble(pivot int, s string) string {
	start := 2*pivot + 2
	bytes := make([]byte, len(s)-start)
	for i := start; i < len(s); i++ {
		bytes[len(s)-1-i] = s[i]
	}
	return string(bytes) + s
}

func checkSingle(pivot int, s string) bool {
	for i, j := pivot-1, pivot+1; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func checkDouble(pivot int, s string) bool {
	if s[pivot] != s[pivot+1] {
		return false
	}
	for i, j := pivot-1, pivot+2; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func test(s string, expected string) {
	actual := shortestPalindrome(s)
	if actual != expected {
		fmt.Printf("%s expected %s, actual %s", s, expected, actual)
	}
}

func main() {
	test("", "")
	test("a", "a")
	test("abcba", "abcba")
	test("abccba", "abccba")
	test("abcbad", "dabcbad")
	test("abbaabcd", "dcbaabbaabcd")
	test("cabbacd", "dcabbacd")
	test("abcbadf", "fdabcbadf")
}
