package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {

	nums := map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	plus := '+'
	minus := '-'
	dot := '.'
	e := 'e'

	s = strings.TrimSpace(s)
	if len(s) == 0 || s == "." {
		return false
	}

	isDotExist := false
	isPlusMinusExist := false
	isNumExist := false
	isFirstDotExist := false
	isEExist := false

	strArr := []rune(s)
	last := strArr[len(strArr)-1]

	if !nums[last] && last != dot {
		return false
	}
	if strArr[0] == dot {
		isDotExist = true
		isFirstDotExist = true
	} else if nums[strArr[0]] {
		isNumExist = true
	} else if strArr[0] == plus || strArr[0] == minus {
		isPlusMinusExist = true
	} else {
		return false
	}

	for i := 1; i < len(s); i++ {
		if strArr[i] == dot {
			if isDotExist || isEExist || (!isNumExist && !isPlusMinusExist) {
				return false
			} else {
				isDotExist = true
			}
		} else if strArr[i] == e {
			if isEExist || (isFirstDotExist && !isNumExist) || strArr[i-1] == minus || strArr[i-1] == plus {
				return false
			} else {
				isEExist = true
			}
		} else if strArr[i] == plus || strArr[i] == minus {
			isPlusMinusExist = true
			if strArr[i-1] != e {
				return false
			}
		} else if nums[strArr[i]] {
			isNumExist = true
		} else {
			return false
		}

	}

	if last == dot {
		return !isEExist && isNumExist
	}

	return true
}

func test(str string, expect bool) {
	if isNumber(str) != expect {
		fmt.Println(str)
	}
}

func play() {
	const s = "日本语"
	const x = '日'
	strArr := []rune(s)
	if strArr[0] == x {
		fmt.Println(1)
	}
	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}
	for _, v := range s {
		fmt.Printf("%#U\n", v)
	}
}

func main() {
	play()
	test("0", true)
	test("3.", true)
	test(" 0.1 ", true)
	test("abc", false)
	test("1 a", false)
	test("2e10", true)
	test(" -90e3   ", true)
	test(" 1e", false)
	test("e3", false)
	test(" 6e-1", true)
	test(" 99e2.5 ", false)
	test("53.5e93", true)
	test(" --6 ", false)
	test("-+3", false)
	test("95a54e53", false)
	test(".e1", false)
	test("-.", false)
	test("-1.", true)
	test("+.8", true)
	test(" 005047e+6", true)
	test("123", true)
	test(" 123 ", true)
	test("0", true)
	test("0123", true) //Cannot agree
	test("00", true)   //Cannot agree
	test("-10", true)
	test("-0", true)
	test("123.5", true)
	test("123.000000", true)
	test("-500.777", true)
	test("0.0000001", true)
	test("0.00000", true)
	test("0.", true)   //Cannot be more disagree!!!
	test("00.5", true) //Strongly cannot agree
	test("123e1", true)
	test("1.23e10", true)
	test("0.5e-10", true)
	test("1.0e4.5", false)
	test("0.5e04", true)
	test("12 3", false)
	test("1a3", false)
	test("", false)
	test("     ", false)
	test(".1", true) //Ok, if you say so
	test(".", false)
	test("2e0", true) //Really?!
	test("+.8", true)
	test(".2e81", true) //Really?!
	test("-e81", false) //Really?!
}
