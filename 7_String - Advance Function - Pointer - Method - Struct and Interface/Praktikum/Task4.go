package main

import (
	"fmt"
)

func minmax (num...*int) (min int,max int){
	min=*num[0]
	max=*num[0]
    for _,i := range num{
        if *i > max {
            max = *i
		}else if *i < min {
            min = *i
		}
    }
	return 
}



func main(){
	var a1,a2,a3,a4,a5,a6,min,max int
	fmt.Print("Masukkan a1 : ")
	fmt.Scan(&a1)
	fmt.Print("Masukkan a2 : ")
	fmt.Scan(&a2)
	fmt.Print("Masukkan a3 : ")
	fmt.Scan(&a3)
	fmt.Print("Masukkan a4 : ")
	fmt.Scan(&a4)
	fmt.Print("Masukkan a5 : ")
	fmt.Scan(&a5)
	fmt.Print("Masukkan a6 : ")	
	fmt.Scan(&a6)

	min,max =minmax(&a1,&a2,&a3,&a4,&a5,&a6)
	fmt.Println("nilai min : ",min)
	fmt.Println("nilai max : ",max)
}

