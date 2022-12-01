// program ini adalah program untuk menghitung akar pangkat 3
package main

import "fmt"

func main() {
	var a, b int
	var habis bool

	fmt.Println("Masukkan 2 bilangan")
	fmt.Scan(&a, &b)

	habis = b%(a*a*a) == 0

	fmt.Println("Apakah akar pangkat 3? ", habis)

}
