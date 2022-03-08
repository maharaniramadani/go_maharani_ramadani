package main

import (
	"fmt"
)

func selec(arr []int){
	for i :=0;i<len(arr);i++{
		mindx := i
		for j:=i+1;j<len(arr);j++{
			if arr[j]<arr[mindx] {
				mindx=j
				
			}
		}
		arr[i], arr[mindx]=arr[mindx],arr[i]
	}
}

func maxbuy(money int,productprice []int){
	selec(productprice)
	maxbrg:=0
	for _,productprice := range productprice{
		if money >= productprice{
			maxbrg +=1
			money -= productprice
		}
	}
	fmt.Println(maxbrg)
}

func main(){
	maxbuy(50000,[]int{25000,25000,10000,14000})
	maxbuy(30000,[]int{15000,12000,5000,3000,10000}) 
	maxbuy(10000,[]int{2000,2000,3000,1000,2000,10000}) 
	maxbuy(4000,[]int{7500,1500,2000,3000})                        // 3
	maxbuy(0,[]int{10000,3000})
}