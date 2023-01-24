//program untuk menentukan jumlah dari beberapa tipe kendaraan 
package main

import "fmt"

func main() {
	var car string
	var TipeA, TipeB, TipeC int

	fmt.Println("Masukkan Tipe Kendaraan : ")
	fmt.Scan(&car)
	for car == "A" || car == "B" || car == "C" {
		if car == "A" {
			TipeA += 1
		} else if car == "B" {
			TipeB += 1
		} else {
			TipeC += 1
		}
		fmt.Scan(&car)
	}
	fmt.Println("Tipe A = ", TipeA)
	fmt.Println("Tipe B = ", TipeB)
	fmt.Println("Tipe C = ", TipeC)
}
