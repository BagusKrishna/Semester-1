// program ini adalah program untuk menentukan apa saja faktor dari bilangan yang diinputkan
package main

import "fmt"

func main() {
	var i, n int
	var apa bool
	fmt.Println("Bilangannya kak? ")
	fmt.Scan(&n)

	fmt.Println()
	fmt.Println("faktornya adalah = ")
	for i = 1; i <= n; i++ {

		apa = n%i == 0
		fmt.Println(i, apa)
	}

}
