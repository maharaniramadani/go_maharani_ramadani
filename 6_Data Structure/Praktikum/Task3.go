// task3

package main

import "fmt"

func pair(arr []int, target int) {
	mapp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		mapp[arr[i]] = i
		index1, item2 := mapp[target-arr[i]]
		index2, item1 := mapp[arr[i]]
		if item1 && item2 {
			fmt.Printf("pasangan index: %d, %d", index1, index2)
			return
		}
	}
	fmt.Println("tidak ada pasangan index")
}

func main() {
	pair([]int{1,2,3,4,6},6)
	pair([]int{2,3,9,11},11)
	pair([]int{1,2,5,7},12)
	pair([]int{1,4,6,8},10)
	pair([]int{1,5,6,7},6)
}