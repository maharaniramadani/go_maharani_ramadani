package main

import "fmt"

func caesar(offset int, input string)string{
    var out string
    var tem rune
    for _,i := range input {
        tem = ((i +rune(offset)  - 97 )%26)+97
        out +=string(tem)
    }
    return out
}


func main(){
	fmt.Println(caesar(3,"abc"))
	fmt.Println(caesar(2,"alta"))
}
