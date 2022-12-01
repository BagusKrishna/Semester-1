// program ini adalah program untuk memasukkan password yang benar dan berapa banyak percobaannya
package main

import "fmt"

func main() {
	var user, pw string
	var i int

	fmt.Println("Masukkan username dan password")
	i = 0
	for user != "admin" || pw != "admin" {
		fmt.Scan(&user, &pw)
		i++
	}
	fmt.Println("Percobaan sebanyak = ", i)
	fmt.Println("Login berhasil")
}
