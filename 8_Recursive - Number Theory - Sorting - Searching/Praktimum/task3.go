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

func primaSegiEmpat(high,wide,start int){
	jum :=0
	start++
	var i,j int
	for i<wide{
		if angkaprima(start) {
			fmt.Printf("%d\t",start)
			jum +=start
			
			j++
		}
		if high==j {
			fmt.Println("")
			j=0
			i++
			
		}
		start++
	}
	fmt.Println(jum)
}

func main (){
	primaSegiEmpat(2,3,13)
	fmt.Println("====================================")
	primaSegiEmpat(5,2,1)
}