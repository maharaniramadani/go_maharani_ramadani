package main

import (
	"fmt"
	
)

func Pangkat(base, pangkat int) int {
	var i, result int
	result = 1
	for i = 0; i < pangkat; i++ {
			result *= base
	}
	return result
}

func main() {
	fmt.Println(Pangkat(2, 10))
	// fmt.Println(Pangkat(5, 3))
	// fmt.Println(Pangkat(10, 2))
	// fmt.Println(Pangkat(2, 5))
	// fmt.Println(Pangkat(7, 3))
}