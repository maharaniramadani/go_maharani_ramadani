// 4. problem bilangan prima
package main

import (
	"fmt"
)

func angkaprima (num int){
	prima := true  
	if num == 0 || num == 1 {  
		fmt.Printf(" %d adalah bukan bilangan prima\n", num)  
		} else {  
			for i := 2; i <= num/2; i++ {  
				if num%i == 0 {  
					fmt.Printf(" %d adalah bukan bilangan prima\n", num)  
					prima = false  
					break  
				}  
			}  
			
			if prima == true {  
				fmt.Printf(" %d adalah bilangan prima\n", num)  
			}  
		}  
}

func main(){
	angkaprima(11)
	angkaprima(13)
	angkaprima(17)
	angkaprima(20)
	angkaprima(35)
}