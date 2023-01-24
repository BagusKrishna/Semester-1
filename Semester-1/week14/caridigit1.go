// program untuk mencari suatu digit x di suatu bilangan tertentu
package main

import "fmt"

func main() {
	var x, n, digit, simpan int
	var benar bool

	fmt.Println("Masukkan digit : ")
	fmt.Scan(&x)
	fmt.Println("Masukkan bilangan : ")
	fmt.Scan(&n)

	if x > n {
		fmt.Println(!benar)
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
		fmt.Println(benar)
	} else {
		fmt.Println(benar)
	}
}
