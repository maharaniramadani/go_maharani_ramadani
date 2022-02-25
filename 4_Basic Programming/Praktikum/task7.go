//problem tiangel

package main

import "fmt"

func main() {
	n:=5
	for i :=0; i < n ;i++{
		for j := 1; j < n-i ; j++{
			fmt.Print(" ") 
		}
		for k:=0; k<=i ;k++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
