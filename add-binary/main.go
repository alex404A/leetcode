package main

import (
	"bytes"
	"fmt"
)

type returnVal struct {
	v         byte
	isCarried bool
}

func addBinary(a string, b string) string {
	var (
		smaller string
		bigger  string
	)
	if len(a) < len(b) {
		bigger = b
		smaller = a
		smaller = supplement(a, len(bigger))
	} else {
		bigger = a
		smaller = b
		smaller = supplement(b, len(bigger))
	}
	isCarried := false
	result := make([]byte, len(bigger)+1)
	for i := len(bigger) - 1; i >= 0; i-- {
		val := cal(bigger[i], smaller[i], isCarried)
		result[i+1] = val.v
		isCarried = val.isCarried
	}
	if isCarried {
		result[0] = '1'
		return string(result)
	} else {
		return string(result[1:])
	}
}

func supplement(x string, biggerLen int) string {
	buf := bytes.Buffer{}
	for i := len(x); i < biggerLen; i++ {
		buf.WriteString("0")
	}
	buf.WriteString(x)
	return buf.String()
}

func copy2Target(biggerPtr *string, targetPtr *[]byte, smallerLen int) {
	bigger := *biggerPtr
	target := *targetPtr
	for i := len(bigger) - smallerLen - 1; i >= 0; i-- {
		target[i+1] = bigger[i]
	}
}

func cal(x byte, y byte, isCarried bool) returnVal {
	if x == '1' && y == '1' {
		if isCarried {
			return returnVal{'1', true}
		} else {
			return returnVal{'0', true}
		}
	} else if x == '1' || y == '1' {
		if isCarried {
			return returnVal{'0', true}
		} else {
			return returnVal{'1', false}
		}
	} else {
		if isCarried {
			return returnVal{'1', false}
		} else {
			return returnVal{'0', false}
		}
	}
}

func test(a string, b string, expected string) {
	actual := addBinary(a, b)
	if actual != expected {
		fmt.Printf("Adding %s and %s does not equals to %s\n", a, b, expected)
		fmt.Printf("%s is actual val\n", actual)
	}
}

func main() {
	test("1010", "1011", "10101")
	test("1010", "1", "1011")
	test("1111", "1", "10000")
}
