package main
import "fmt"

var memo map[int]int=map[int]int{}

func fibo(n int)int{
	if nilai,sdhHtg:=memo[n];sdhHtg {
		return nilai
	}

	if n<=1{
		memo[n]=n
	}else{
		memo[n]=fibo(n-1)+fibo(n-2)
	}
	return memo[n]
}

var memoTd map[int]int=map[int]int{}

func fibotd(n int)int{
	nilai,sdhHtg := memoTd[n]
	if sdhHtg{
		return nilai
	}

	if  n<= 1{
		memoTd[n]=n
	}else{
		memoTd[n]= fibotd(n-1)+fibotd(n-1)
	} 
		return memoTd[n]
	
}

// var memoBU map[int]int=map[int]int{}

func fiboBU(n int)int{
	var memoBU map[int]int=map[int]int{}
	for i := 0;i<=n;i++{
		if n <=i {
			memoBU[i]=i
			
		}else{
			memoBU[i]=memoBU[i-1]+memoBU[i-2]
		}
	}
	return memoBU[n]
}

func main(){
	// fmt.Println(fibo(0)) 
	// fmt.Println(fibo(1)) // 1
	// fmt.Println(fibo(2)) // 1
	// fmt.Println(fibo(3)) // 2
	// fmt.Println(fibo(5)) // 5
	// fmt.Println(fibo(6)) // 8
	// fmt.Println(fibo(7)) // 13
	// fmt.Println(fibo(9)) // 34
	// fmt.Println(fibo(10)) // 55
	// // fmt.Println(fiboCollect)
	// fmt.Println(fibo(50)) // 55

	// fmt.Println(fibotd(0)) 
	// fmt.Println(fibotd(1)) // 1
	// fmt.Println(fibotd(2)) // 1
	// fmt.Println(fibotd(3)) // 2
	// fmt.Println(fibotd(5)) // 5
	// fmt.Println(fibotd(6)) // 8
	// fmt.Println(fibotd(7)) // 13
	// fmt.Println(fibotd(9)) // 34
	// fmt.Println(fibotd(10)) // 55
	// // fmt.Println(fiboCollect)
	// fmt.Println(fibotd(50)) // 55

	fmt.Println(fiboBU(0)) 
	fmt.Println(fiboBU(1)) // 1
	fmt.Println(fiboBU(2)) // 1
	fmt.Println(fiboBU(3)) // 2
	fmt.Println(fiboBU(5)) // 5
	fmt.Println(fiboBU(6)) // 8
	fmt.Println(fiboBU(7)) // 13
	fmt.Println(fiboBU(9)) // 34
	fmt.Println(fiboBU(10)) // 55
	// fmt.Println(fiboCollect)
	fmt.Println(fiboBU(50)) // 55
}