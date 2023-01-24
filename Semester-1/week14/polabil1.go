//program untuk membuat pola sesuai urutan angka secara horizontal
package main

import "fmt"

func main() {
	var i, j, n int

	fmt.Scan(&n)
	for j = 1; j <= n; j++ {
		for i = 1; i <= n; i++ {

			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
