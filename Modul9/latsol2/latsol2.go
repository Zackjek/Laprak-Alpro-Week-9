package main

import "fmt"

func main() {
	var a int
	var hasil string
	fmt.Scan(&a)
	hasil = "bukan"
	if a < 0 && a%2 == 0{
		hasil = "genap negatif"
	}
	fmt.Print(hasil)
}