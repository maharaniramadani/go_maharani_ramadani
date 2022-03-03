package main

import (
	"fmt"
)

func ArrayMerge(arrayA, ArrayB []string) []string {
  
  mapp := make(map[string]int)
	newArray := append(arrayA,ArrayB...)
	result := make([]string,0)
	for _, i := range newArray {
		mapp[i] = 1
	}
  for j := range mapp {
    result = append(result,j)
  }
  return result
}


func main() {
  fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"king","eddie", "steve", "geese"}))
  // ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
  fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
  // ["sergei", "jin", "steve", "bryan"]
  fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
  // ["alisa", "yoshimitsu", "devil jin", "law"]
  fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
  // ["devil jin", "sergei"]
  fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
  // ["hwoarang"]
  fmt.Println(ArrayMerge([]string{}, []string{}))
  // []
}
