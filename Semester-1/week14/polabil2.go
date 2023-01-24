//program untuk membuat pola sesuai urutan angka secara vertikal
package main

import "fmt"

func main() {
	var i, n, kurang, simpen int
	var j int

	fmt.Scan(&n)

	kurang = n - 1
	for j = 1; j <= n; j++ {
		for i = 1; i <= n; i++ {

			simpen = n - kurang
			fmt.Print(simpen, " ")
		}
		fmt.Println()
		kurang = kurang - 1
	}
}
