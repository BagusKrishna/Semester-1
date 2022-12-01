// program ini adalah program untuk membalikkan digit yang dimasukkan kemudian menjumlahkannya
package main

import "fmt"

func main() {
	var angka, digit, jumlah int

	fmt.Println("Masukkan angka")
	fmt.Scan(&angka)
	jumlah = 0

	for angka > 0 {
		digit = angka % 10

		jumlah = jumlah + digit
		angka = angka / 10
		fmt.Println("digit = ", digit)

	}

	fmt.Println("Hasil penjumlahan digit = ", jumlah)
}
