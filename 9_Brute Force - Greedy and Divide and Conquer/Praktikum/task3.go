package main
import "fmt"

func selec(arr *[]int){
	for i :=0;i<len(*arr);i++{
		for j:=i+1;j<len(*arr);j++{
			if (*arr)[j]< (*arr)[i] {
				temp := (*arr)[i]
				(*arr)[i]=(*arr)[j]
				(*arr)[j] = temp
				
			}
		}
	
	}
}

func dragon(dragonHead, knight []int){
	if len(knight)<len(dragonHead){
		fmt.Println("knight fall")
		return
	}
	
	selec(&dragonHead)
	selec(&knight)

	count:= 0
	result := 0

	for i :=0; i<len(dragonHead);i++{
		for j :=0; j <len(knight);j++{
			if dragonHead[i]<= knight[j]{
				result +=knight[j]
				count++
				break
			}
		}
	}
	if count !=len(dragonHead){
		fmt.Println("knight fall")

	}else{
		fmt.Println(result)
	}
}

func main(){
	dragon([]int{5,4},[]int{7,8,4})
	dragon([]int{5,10},[]int{5})
	dragon([]int{7,2},[]int{4,3,1,2})
	dragon([]int{7,2},[]int{2,1,8,5})
}