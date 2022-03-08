package main
import "fmt"

func playdom(card [][]int,deck []int)interface{}{
	total:= 0
	result := make([]int,2)
	for _,valcard :=range card{
		for _, valdeck := range deck {
			if valcard[0]==valdeck||valcard[1]==valdeck {
				if valcard[0]+valcard[1]>total {
					total = valcard[0]+valcard[1]
					result[0], result[1]=valcard[0], valcard[1]
					
				}
				
			}
		}
	}
	if total !=0 {
		return result
		
	}else{
		return "tutup kartu"
	}
}

func main(){
	fmt.Println(playdom([][]int{{6,5},{3,4},{2,3},{3,3}},[]int{4,3}))
	fmt.Println(playdom([][]int{{6,5},{3,3},{3,4},{2,1}},[]int{3,6}))
	fmt.Println(playdom([][]int{{6,6},{2,4},{3,6}},[]int{5,1}))
}