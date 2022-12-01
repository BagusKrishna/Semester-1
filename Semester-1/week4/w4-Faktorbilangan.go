//program ini adalah program menghitung faktor suatu bilangan

package main

import "fmt"

func main() {
	var a, b int
	var habis bool

	fmt.Println("Masukkan 2 bilangan")
	fmt.Scan(&a, &b)

	habis = b%a == 0

	fmt.Println("Apakah habis ? ", habis)

}
