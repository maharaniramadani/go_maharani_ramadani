package main

import (
	"fmt" 
)

func munculsekali (num  []int) int {
    result := 0
	for i := 0; i < len(num); i++ {
		result ^= num[i]
	}
	return result
}

func main(){
	fmt.Println(munculsekali ([]int{1,2,3,4,1,2,3}))
	fmt.Println(munculsekali ([]int{7,6,5,2,3,7,5,2}))
	fmt.Println(munculsekali ([]int{1,2,3,4,5}))
	fmt.Println(munculsekali ([]int{1,1,2,2,3,3,4,4,5,5}))
	fmt.Println(munculsekali ([]int{0,8,7,2,5,0,4}))
}