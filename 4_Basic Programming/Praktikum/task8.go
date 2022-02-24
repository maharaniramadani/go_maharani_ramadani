
// tabel perkalian
package main

import "fmt"

func perkalian(n int) {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            kali := j*i
            fmt.Scan(&kali)
            
        } 
        
    }
    
}

func main(){
    perkalian(9)
}