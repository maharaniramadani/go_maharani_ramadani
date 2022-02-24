
// tabel perkalian
package main

import "fmt"

func perkalian(n int) {
    var kali int
    kali=1
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            kali = i*j
            fmt.Print("",kali,"\t")
        }   
    fmt.Println("")
    }
}

func main(){
    perkalian(9)
}