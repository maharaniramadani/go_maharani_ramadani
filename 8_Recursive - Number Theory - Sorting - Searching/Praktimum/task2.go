package main
import "fmt"

func fibo(num int)int{
	if num ==0{
		return 0
	}else if num==1{
		return 1
	}
	return fibo(num-1)+fibo(num-2)
}

func main (){
	fmt.Println(fibo(0))
	fmt.Println(fibo(2))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
	fmt.Println(fibo(12))
}