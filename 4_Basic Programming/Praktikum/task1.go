package main

import "fmt"

func main(){
	// TODO: 1. luas permukaan tabung
	// r := 4.0
	// t := 20.0
	// p := 3.14
	// Luas := 2.0*p*r*(r+t)
	// fmt.Println("Luas permukaan tabung",Luas)

	// TODO:luas dengan inputan scan

	var t,r float32
	var phi float32 = 3.14

	fmt.Println("Menghitung Luas Permukaan Tabung")
	fmt.Println("==================================")
	fmt.Print("Masukkan tinggi tabung : ")
	fmt.Scan(&t)

	fmt.Print("jari-jari : ")
	fmt.Scan(&r)

	luas := 2*phi*r*(r+ t)
	fmt.Println("\nLuas permukaan tabung: ",luas)
}