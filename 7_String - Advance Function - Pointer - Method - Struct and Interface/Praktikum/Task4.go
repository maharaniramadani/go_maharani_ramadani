package main

import (
	"fmt"
)

func minmax (num...*int) (min int,max int){
	var k int 
    for k=1;k < len(num); k++ {
        if *(num[k]) > max {
            max = *(num[k])
		}else if *(num[k]) < min {
            min = *(num[k])
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

