// program untuk mengetahui beberapa hal diantaranya :
// kartu apabila pembeli bersedia lalu memperoleh diskon
// Diskon 10% apabila belanja minimal 100k.
// cashback apabila belanja minimal 200k dan punya kartu
// kemudian total bayaran yang harus di bayar.

package main

import "fmt"

func main() {
	var belanja int
	var card string
	var kartu bool
	var diskon bool
	var cashback bool

	fmt.Print("Total belanja: ")
	fmt.Scan(&belanja)
	fmt.Print("New Membership Card? ")
	fmt.Scan(&card)

	kartu = card == "true"
	diskon = belanja > 100000
	cashback = belanja > 200000 && card == "true"

	fmt.Println("Kartu: ", kartu)
	fmt.Println("Diskon: ", diskon)
	fmt.Println("Cashback: ", cashback)

}
