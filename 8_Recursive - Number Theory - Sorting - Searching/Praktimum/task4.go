package main
import "fmt"

func maxseq(arr []int)int{
	result := 0
	for i,_:= range arr{
		sum:=0
		for _,value := range arr[i:]{
			sum += value
			if sum> result{
				result=sum
			}
		}
	}
	return result
}

func main(){
	fmt.Println(maxseq([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxseq([]int{-2,-5,6,-2,-3,1,5,-6})) 
	fmt.Println(maxseq([]int{-2,-3,4,-1,-2,1,5,-3})) 
	fmt.Println(maxseq([]int{-2,-5,6,-2,-3,1,6,-6})) 
	fmt.Println(maxseq([]int{-2,-5,6,2,-3,1,6,-6}))
}