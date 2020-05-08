package main

import "fmt"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sum := rectangleArea(A, B, C, D) + rectangleArea(E, F, G, H)
	return sum - intersection(A, B, C, D, E, F, G, H)
}

func rectangleArea(ldx int, ldy int, rux int, ruy int) int {
	return (rux - ldx) * (ruy - ldy)
}

func intersection(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if C <= E || G <= A {
		return 0
	}
	if D <= F || H <= B {
		return 0
	}
	height := 0
	width := 0
	if A <= E {
		if C <= G {
			width = C - E
		} else {
			width = G - E
		}
	} else {
		if C <= G {
			width = C - A
		} else {
			width = G - A
		}
	}
	if D >= H {
		if B >= F {
			height = H - B
		} else {
			height = H - F
		}
	} else {
		if B >= F {
			height = D - B
		} else {
			height = D - F
		}
	}
	return height * width
}

func main() {
	result := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	fmt.Println(result)
}
