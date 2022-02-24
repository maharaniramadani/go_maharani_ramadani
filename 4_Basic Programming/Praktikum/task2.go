//TODO: 2. grade nilai
package main

import "fmt"

func main(){

	var nilai int
	var nama string

	fmt.Println("Aplikasi Input nilai")
	fmt.Println("====================")
	fmt.Print("Nama siswa: ")
	fmt.Scanln(&nama)
	fmt.Print("Nilai siswa: ")
	fmt.Scan(&nilai)

	// var nilai int = 75

	if nilai > 80 && nilai < 100{
		fmt.Println("A")
	}else if nilai < 79 && nilai > 65 {
		fmt.Println("B")
	}else if nilai < 64 && nilai > 50 {
		fmt.Println("C")
	}else if nilai < 49 && nilai > 35 {
		fmt.Println("D")
	}else {
		fmt.Println("E")
	}

}