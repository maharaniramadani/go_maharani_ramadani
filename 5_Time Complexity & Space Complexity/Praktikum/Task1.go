package main

import (
	"fmt"
	"math/big"
	"time"
)

func angkaprima (num int64){
    if big.NewInt(num).ProbablyPrime(0) {
		fmt.Println(num, " True")
		time.Sleep(1000 * time.Millisecond)
	} else {
		fmt.Println(num, " False")
		time.Sleep(1000 * time.Millisecond)
	}
}
	
func main(){
	start := time.Now()
	angkaprima(1000000007)
	fmt.Println("lama waktu:", time.Since(start))
	angkaprima(13)
	fmt.Println("lama waktu:", time.Since(start))
	angkaprima(17)
	fmt.Println("lama waktu:", time.Since(start))
    angkaprima(20)
	fmt.Println("lama waktu:", time.Since(start))
	angkaprima(35)
    fmt.Println("lama waktu:", time.Since(start))
}	