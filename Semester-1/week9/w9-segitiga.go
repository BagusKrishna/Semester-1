//program untuk mengetahui apakah bilangan yang dimasukkan dapat memnbentuk segitiga atau tidak.

package main

import "fmt"

func main() {
	var a, b, c int
	var segitiga bool

	fmt.Println("bilangan?")
	fmt.Scan(&a, &b, &c)

	segitiga = (a >= 0 && b >= 0 && c >= 0) && (a >= b+c || b >= a+c || c >= a+b) || a == b || a == c || b == c
	fmt.Println("Apakah segitiga?? ", segitiga)
}
