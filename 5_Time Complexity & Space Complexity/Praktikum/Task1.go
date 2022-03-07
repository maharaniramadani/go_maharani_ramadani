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

 func main(){
	fmt.Println(angkaprima(1000000007))
	fmt.Println(angkaprima(13))
	fmt.Println(angkaprima(17))
    fmt.Println(angkaprima(20))
	fmt.Println(angkaprima(35))
}	