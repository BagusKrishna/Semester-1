//Program untuk menentukan tiga digit nilai yang terdapat pada suatu bilangan bulat positif x

package main

import "fmt"

func main() {
	var a, belakang, tengah, depan int

	fmt.Println("Masukkan 1 bilangan (maksimal 3 digit)")
	fmt.Scan(&a)

	belakang = a % 10
	depan = a / 100
	tengah = (a / 10) % 10

	fmt.Print("angka kamu = ", depan, tengah, belakang)
}
