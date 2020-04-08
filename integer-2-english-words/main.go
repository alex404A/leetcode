package main

import (
	"fmt"
	"strconv"
)

func help(raw string, str string) string {
	if raw == "" {
		return str
	} else {
		return str + " " + raw
	}
}

func deal3digits(str string) string {
	lessThan20 := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	num, _ := strconv.Atoi(str)
	if num == 0 {
		return ""
	}
	if num < 20 {
		return lessThan20[num-1]
	}
	result := ""
	unit, _ := strconv.Atoi(str[len(str)-1 : len(str)])
	ten, _ := strconv.Atoi(str[len(str)-2 : len(str)-1])
	if ten <= 1 {
		if ten != 0 || unit != 0 {
			two, _ := strconv.Atoi(str[len(str)-2 : len(str)])
			result += lessThan20[two-1]
		}
	} else {
		if unit != 0 {
			result += lessThan20[unit-1]
		}
		result = help(result, tens[ten-2])
	}
	if num < 100 {
		return result
	} else {
		hundred, _ := strconv.Atoi(str[0:1])
		result = help(result, lessThan20[hundred-1]+" Hundred")
	}
	return result
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	thousands := []string{"Thousand", "Million", "Billion"}
	result := ""
	clusters := split(num)
	for i := 0; i < len(clusters); i++ {
		part := deal3digits(clusters[i])
		if part == "" {
			continue
		} else {
			if i > 0 {
				part = part + " " + thousands[i-1]
			}
			if result == "" {
				result = part
			} else {
				result = part + " " + result
			}
		}
	}
	return result
}

func split(num int) []string {
	str := strconv.Itoa(num)
	start := ifthenelse(len(str)-3 >= 0, len(str)-3, 0)
	end := ifthenelse(start+3 <= len(str), start+3, len(str))
	result := []string{}
	for ; start >= 0; start -= 3 {
		result = append(result, str[start:end])
		end -= 3
	}
	if end > 0 {
		result = append(result, str[0:end])
	}
	return result
}

func ifthenelse(cond bool, i int, j int) int {
	if cond {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(numberToWords(101))
	fmt.Println(numberToWords(2))
	fmt.Println(numberToWords(12))
	fmt.Println(numberToWords(32))
	fmt.Println(numberToWords(0))
	fmt.Println(numberToWords(1))
	fmt.Println(numberToWords(11))
	fmt.Println(numberToWords(21))
	fmt.Println(numberToWords(100))
	fmt.Println(numberToWords(111))
	fmt.Println(numberToWords(121))
	fmt.Println(numberToWords(619))
	fmt.Println(numberToWords(621))
	fmt.Println(numberToWords(1621))
	fmt.Println(numberToWords(18621))
	fmt.Println(numberToWords(21621))
	fmt.Println(numberToWords(100621))
	fmt.Println(numberToWords(111621))
	fmt.Println(numberToWords(121621))
	fmt.Println(numberToWords(121000))
	fmt.Println(numberToWords(1000000))
}
