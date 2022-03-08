package main
import "fmt"

func minmax (arr []int) string{
	min:= arr[0]
	max:= arr[0]
	indexmax := 0
	indexmin := 0

	for i,value := range arr {
		if value> max {
			max=value
			indexmax=i
			
		}
		if value<min{
			min=value
			indexmin=i
		
		}
	}
	return fmt.Sprintf("max : %d index: %d min: %d index: %d",max,indexmax,min,indexmin)
}

func main(){
	fmt.Println(minmax([]int{5,7,4,-2,-1,8}))
	fmt.Println(minmax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(minmax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(minmax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(minmax([]int{-2, 5, -7, 4, 7, -20}))
}