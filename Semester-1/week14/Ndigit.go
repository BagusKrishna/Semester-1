// program untuk mengetahui digit terbesar dari suatu bilangan yang diinput
package main

import "fmt"

func main() {
	var bil, terbesar int
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&bil)

	terbesar = bil % 10
	bil = bil / 10

	for bil/10 != 0 {
		if bil%10 > terbesar {
			terbesar = bil % 10
		}
		if (bil/10 > terbesar) && (bil/10 < 10) {
			terbesar = bil / 10
		}
		bil = bil / 10
	}
	fmt.Println("digit terbesarnya adalah", terbesar)
}
