package main

import (
	"fmt"
)

func moneycoin(money int)[]int{
	coin := []int{1,10,20,50,100,200,500,1000,2000,5000,10000}
	var res []int
	for i := len(coin)-1;i>=0;i--{
		for coin[i] <= money{
		res = append(res,coin[i])
		money -= coin[i]
		}
	}
	return res
}

func main(){
fmt.Println(moneycoin(123))
fmt.Println(moneycoin(432))
fmt.Println(moneycoin(543))
fmt.Println(moneycoin(7752))
fmt.Println(moneycoin(15321))
}
