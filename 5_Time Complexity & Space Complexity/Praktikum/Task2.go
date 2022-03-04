package main

import (
	"fmt"
	
)

func Pangkat(base, pangkat int) int {
	result := 1
	for pangkat >0{
		if pangkat%2 == 1 {
			result*=base
		}
		pangkat/=2
		base*=base
	}
	return result
}

func main() {
	fmt.Println(Pangkat(2, 10))
	fmt.Println(Pangkat(5, 3))
	fmt.Println(Pangkat(10, 2))
	fmt.Println(Pangkat(2, 5))
	fmt.Println(Pangkat(7, 3))
}