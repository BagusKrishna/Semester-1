package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("masukkan 3 angka bos")
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Print("Ini segitiga sama sisi")
	} else if a == b || b == c || a == c {
		fmt.Print("Ini segitiga sama kaki")
	} else {
		fmt.Print("Ini segitia sembarang")
	}
}
