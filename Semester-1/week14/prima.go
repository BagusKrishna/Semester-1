//Program untuk mengetahui apakah angka yang dimasukkan merupakan bilangan prima atau tidak.

package main

import "fmt"

func main() {
	var bil, i, jumlah int
	fmt.Println("Masukkan bilangan: ")
	fmt.Scan(&bil)

	for i = 1; i <= bil; i++ {
		if bil%i == 0 {
			jumlah += 1
		}
	}

	if jumlah == 2 {
		fmt.Print("prima")
	} else {
		fmt.Print("bukan prima")
	}

}
