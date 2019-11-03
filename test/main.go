package main

import "fmt"

func test(slice []int) {
	slice[0] = 1
	fmt.Println(slice)
}

func main() {
	slice := []int{0, 0}
	test(slice)
	fmt.Println(slice)
}
