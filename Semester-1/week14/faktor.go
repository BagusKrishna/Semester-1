// program untuk mengetahui apa saja faktor dari suatu bilangan yang diinput
package main

import "fmt"

func main() {
	var bil, i int

	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&bil)

	fmt.Println("faktor dari", bil, "adalah: ")
	for i = 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(i, " ")
		}
	}

}
