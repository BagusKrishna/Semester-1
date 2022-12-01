// program untuk menghitung kapasitas tanki yang penuh apabila ditambahkan air
package main

import "fmt"

func main() {
	var tangki, ember, jumlah int
	var penuh bool

	fmt.Println("Berapa kapasitas tanki? ")
	fmt.Scan(&tangki)
	jumlah = 0

	for jumlah <= tangki {
		fmt.Println("Volume air yang akan ditambahkan ?")
		fmt.Scan(&ember)
		jumlah = jumlah + ember
		penuh = jumlah >= tangki
		fmt.Println(penuh)
		if jumlah >= tangki {
			break
		}
	}
}
