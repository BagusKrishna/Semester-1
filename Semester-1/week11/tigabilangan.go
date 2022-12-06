/*
	Program untuk mengurutkan 3 bilangan
	Bilangan diurutkan dari yang terkecil ke terbesar
*/

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Berapa bilanganmu??")
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Print(a, b, c)
	} else if a >= b && a >= c {
		if b >= c {
			fmt.Print(c, b, a)
		} else {
			fmt.Print(b, c, a)
		}
	} else if b >= a && b >= c {
		if a >= c {
			fmt.Print(c, a, b)
		} else {
			fmt.Print(a, c, b)
		}
	} else if c >= a && c >= b {
		if a >= b {
			fmt.Print(b, a, c)
		} else {
			fmt.Print(a, b, c)
		}
	}
}
