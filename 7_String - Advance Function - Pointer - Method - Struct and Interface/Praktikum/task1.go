package main

import (
	"fmt"
)

func compare(a, b string) string {
	new := make([][]int, 1 + len(a))

	for i := 0; i < len(new); i++ {
		new[i] = make([]int, 1 + len(b))
	}

	panjang := 0
	x_panjang := 0

	for x := 1; x < 1 + len(a); x++ {
		for y := 1; y < 1 + len(b); y++ {
			if a[x-1] == b[y-1] {
				new[x][y] = new[x-1][y-1] + 1

			if new[x][y] > panjang {
				panjang = new[x][y]
				x_panjang = x
			}
			}
		}
	}
	return a[x_panjang - panjang: x_panjang]
}
	  
func main(){
	fmt.Println(compare("aka","akasi"))
	fmt.Println(compare("kangguru","guru"))
}

