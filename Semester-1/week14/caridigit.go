// program untuk mencari suatu digit x di suatu bilangan tertentu
package main

import "fmt"

func main() {
	var x, n, digit, simpan, sini, ini int
	var benar bool

	fmt.Println("Masukkan digit : ")
	fmt.Scan(&x)
	fmt.Println("Masukkan bilangan : ")
	fmt.Scan(&n)

	ini = x
	sini = n
	if x > n {
		fmt.Println(ini, "GAK ADA pada bilangan", sini)
	} else {
		digit = 1
		simpan = x
		for simpan > 0 {
			digit = digit * 10
			simpan = simpan / 10
		}
		for n != 0 {
			if n%digit == x {
				benar = true
			}
			n = n / 10
		}
	}
	if benar {
		fmt.Println(ini, "ADA pada bilangan", sini)
	} else {
		fmt.Println(ini, "GAK ADA pada bilangan", sini)
	}
}