package main

import (
	"fmt"
	"math"
)

 func prime(num int)bool{
	for i := 2; i <= int(math.Floor(float64(num) / 2)); i++ {
        if num%i == 0 {
            return false
        }
    }
    return num > 1
 }

 func main(){
	 fmt.Println(prime(1000000007))
 }