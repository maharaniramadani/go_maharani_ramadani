// TODO: 3. faktor bilangan
package main
import "fmt"

func main() {

    var angka, i int
    fmt.Print("Masukkan Angka faktor: ")
    fmt.Scan(&angka)
    fmt.Println("Faktor dari:",angka)
    for i = 1; i <= angka; i++ {
        if angka%i == 0 {
            fmt.Println(i)
        }
    }
}