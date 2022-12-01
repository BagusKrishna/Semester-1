// program ini adalah program untuk menghitung faktorial dari suatu bilangan
package main

import "fmt"

func main() {
	var i, n, total int
	fmt.Println("Bilangannya kak? ")
	fmt.Scan(&n)

	total = 1
	for i = 1; i <= n; i++ {
		total = total * i
	}
	fmt.Println("hasilnya adalah ", total)
}
