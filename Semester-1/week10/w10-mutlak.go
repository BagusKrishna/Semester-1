package main

import "fmt"

func main() {
	var n int

	fmt.Println("masukkan angkamu bro")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Print("Hasil = ", n*-1)
	} else {
		fmt.Print("Hasil = ", n)
	}
}
