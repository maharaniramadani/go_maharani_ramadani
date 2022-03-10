package main
import "fmt"

func Binary(arr []int,x int){
	res :=-1
	kiri :=0
	kanan:=len(arr)-1
	for kiri <= kanan && res == -1{
		mid := (kiri+kanan)/2

		if x<arr[mid] {
			kanan = mid-1
		}else if x > arr[mid] {
			kiri =mid+1
		}else{
			res =mid
		}
	}
	fmt.Println(res)

}

func main(){
	Binary([]int{1,2,3,5,5,6,7},3)
	Binary([]int{1,2,3,5,6,8,10},5)
	Binary([]int{12,15,15,19,24,31,53,59,60},53)
	Binary([]int{12,15,25,29,24,31,53,59,60},100)
}