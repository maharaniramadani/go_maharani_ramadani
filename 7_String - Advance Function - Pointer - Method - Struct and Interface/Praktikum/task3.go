package main

import (
	"fmt"
)

func swap(a,b*int){
	pointer:= *a
	*a = *b
	*b = pointer
}

func main(){
	a:=10
	b:=20
	swap(&a,&b)
	fmt.Println("setelah di swap menjadi : ",a,b)
}