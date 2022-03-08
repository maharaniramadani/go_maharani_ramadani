package main
import "fmt"

type pair struct{
	name string
	count int
}

func Sort(arr []pair){
	for i :=0;i<len(arr);i++{
		mindx := i
		for j:=i+1;j<len(arr);j++{
			if arr[j].count < arr[mindx].count {
				mindx=j
				
			}
		}
		arr[i], arr[mindx]=arr[mindx],arr[i]
	}
}

func appear(item []string)[]pair{
	mapp:= map[string]int{}
	for _,i :=range item{
		k := i
		_,e := mapp[k]
		if e{
			mapp[k]++
		}else{
			mapp[k]=1
		}
	}
	p:=[]pair{}
	for k,v :=range mapp{
		np:=pair{
			name: k,
			count: v,
		}
		p=append(p, np)
	}
	Sort(p)
	return p
}

func main(){
	fmt.Println(appear([]string{"js","js","golang","ruby","ruby","js","js"}))
	fmt.Println(appear([]string{"A","B","B","C","A","A","B","A","D","D"}))
	fmt.Println(appear([]string{"football","basketball","tenis"}))
}