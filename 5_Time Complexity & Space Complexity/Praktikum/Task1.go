package main

import (
	"fmt"
)

 func angkaprima(num int)bool{
	for i := 2; i <= int(num)/2; i++ {
        if num%i == 0 {
            return false
        }
    }
    return num > 1
 }

 func main(){
	fmt.Println(angkaprima(1000000007))
	fmt.Println(angkaprima(13))
	fmt.Println(angkaprima(17))
    fmt.Println(angkaprima(20))
	fmt.Println(angkaprima(35))
}	