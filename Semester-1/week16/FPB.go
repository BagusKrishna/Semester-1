// program untuk mencari faktor persekutuan terbesar dari kedua bilangan.
// Contoh : 
// FAKTOR dari 5 = 1,5
// FAKTOR dari 10 = 1,2,5,10
// Maka Faktor terbesar kedua bilangan tersebut adalah 5
package main

import "fmt"

func main() {
	var a, b, terbesar, i, fterbesar int

	fmt.Scan(&a, &b)
	fterbesar = 0
	if a > b {
		terbesar = a
	} else {
		terbesar = b
	}
	for i = 1; i <= terbesar; i++ {
		if a%i == 0 && b%i == 0 {
			if i > fterbesar{
				fterbesar = i
			}
		}
	}
	fmt.Print(fterbesar)
}
