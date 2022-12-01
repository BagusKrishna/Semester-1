package main

import "fmt"

func main() {
	var a, b, hasil float64

	fmt.Println("Masukkan 2 bilangan")
	fmt.Scan(&a, &b)

	if a > b {
		hasil = a - b
		fmt.Print("Waduh, kamu mengalami penurunan sebesar ")
		fmt.Print(hasil)
	} else if a < b {
		hasil = b - a
		fmt.Print("Asikk, kamu lagi profit sebesar ")
		fmt.Print(hasil)
	} else {
		fmt.Print("sama aja humm")
	}
}
