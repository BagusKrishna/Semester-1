// program yang menampilkan apa saja faktor dari kedua bilangan yang diinput.
package main

import "fmt"

func main() {
	var a, b, terbesar, i int

	fmt.Scan(&a, &b)
	if a > b {
		terbesar = a
	} else {
		terbesar = b
	}
	for i = 1; i <= terbesar; i++ {
		if a%i == 0 && b%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
