package main

import "fmt"

func main() {
	var a, b, c, d, e float64

	fmt.Println("Masukkin 5x mas")
	fmt.Scan(&a, &b, &c, &d, &e)

	if a > b && b > c && c > d && d > e {
		fmt.Print("Ini stabil turun kak")
	} else if e > d && d > c && c > b && b > a {
		fmt.Print("Ini stabil naik kak")
	} else {
		fmt.Print("waduh, gak stabil ini brooo..")
	}
}
