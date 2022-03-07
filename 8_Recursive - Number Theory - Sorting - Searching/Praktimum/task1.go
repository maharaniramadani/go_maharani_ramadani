package main

import (
	"fmt"
	"math"
)

func angkaprima(num int)bool{
	if num< 2 {
		return false
	}
	sqrtnum := int(math.Sqrt(float64(num)))
	for i := 2; i<=sqrtnum;i++{
		if num%i==0 {
			return false
		}
	}
	return true
 }

func primx (num int) int{
	index:=[]int{}

	for i:=2;true;i++{
		if angkaprima(i){
			index=append(index, i)
		}
		if len(index)==num{
			break
		}
	}
	return index[num-1]
}

func main(){
	fmt.Println(primx(1))
	fmt.Println(primx(5))
	fmt.Println(primx(8))
	fmt.Println(primx(9))
	fmt.Println(primx(10))
}