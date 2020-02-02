package main

import "fmt"

func isPalindrome(s string) bool {
	bytes := filter(s)
	if len(bytes) == 0 {
		return true
	}
	return check(bytes)
}

func filter(s string) []byte {
	result := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= 65 && b <= 90 {
			result = append(result, b+32)
		} else if (b >= 97 && b <= 122) || (b >= 48 && b <= 57) {
			result = append(result, b)
		}
	}
	return result
}

func check(bytes []byte) bool {
	length := len(bytes)
	for i := 0; i < len(bytes); i++ {
		head := bytes[i]
		tail := bytes[length-1-i]
		if head != tail {
			return false
		}
	}
	return true
}

func test(s string, expected bool) {
	actual := isPalindrome(s)
	if actual != expected {
		fmt.Printf("%s\n", s)
	}
}

func main() {
	test("A man, a plan, a canal: Panama", true)
	test("race a car", false)
}
