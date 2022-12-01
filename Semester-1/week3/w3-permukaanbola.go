//program untuk menghitung luas permukaan bola

package main

import "fmt"

func main() {
	var r, luas float64

	fmt.Println("Masukkan jari-jari")

	fmt.Scan(&r)

	luas = 4 * r * r * 22 / 7
	fmt.Println("Luas permukaan bola adalah ", luas)

}
