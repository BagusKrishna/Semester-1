/* 
	Program untuk menghitung total uangmu dalam 4 atau 5 minggu
*/
package main

import "fmt"

func main() {
	var uang1, uang2, uang3, uang4, uang5, minggu, total int

	fmt.Println("berapa minggu?")
	fmt.Scan(&minggu)

	if minggu == 4 {
		fmt.Println("Banyak uang? (sesuai jumlah  minggu)")
		fmt.Scan(&uang1, &uang2, &uang3, &uang4)
		total = uang1 + uang2 + uang3 + uang4
		fmt.Println("Total uangmuu = ", total)

	} else if minggu == 5 {
		fmt.Println("Banyak uang? (sesuai jumlah  minggu)")
		fmt.Scan(&uang1, &uang2, &uang3, &uang4, &uang5)
		total = uang1 + uang2 + uang3 + uang4 + uang5
		fmt.Println("Total uangmuu = ", total)
	} else {
		fmt.Println("Minggu mu aneh mas")
	}

	/*
		for i = 1; i <= minggu; i++ {
			fmt.Scan(&uang)
			total = total + uang
		}

		fmt.Println("Total uangmuu = ", total)

	*/
}
