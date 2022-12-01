// program ini adalah program untuk menghitung 3 baris bilangan selanjutnya dari 2 digit bilangan yang diinputkan
package main

import "fmt"

func main() {
	var a, b int
	var pertama, kedua, ketiga int

	fmt.Println("Masukkan 2 bilangan")
	fmt.Scan(&a, &b)

	pertama = a + b
	kedua = pertama + b
	ketiga = pertama + kedua

	fmt.Println("baris bilangan selanjutnya adalah ", pertama, kedua, ketiga)

}
