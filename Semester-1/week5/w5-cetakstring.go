// program ini adalah program untuk mengulang sebuah karakter sebanyak nilai tertentu
package main

import "fmt"

func main() {
	var a, i int
	var kata string

	fmt.Println("Mau diulang berapa kali? ")
	fmt.Scan(&a)
	fmt.Println("Apa yang mau diulang? ")
	fmt.Scan(&kata)

	fmt.Println("Keluaran = ")
	fmt.Println()
	for i = 1; i <= a; i++ {

		fmt.Println(kata)
	}
}
