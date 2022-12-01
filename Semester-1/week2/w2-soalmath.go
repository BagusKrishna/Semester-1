//ini adalah program untuk menghitung suatu fungsi f(X,y) = 1/3x² + 10 y + 7

package main

import "fmt"

func main() {
	var x, y, hasil float64

	fmt.Println("Masukkan x dan y")
	fmt.Scan(&x, &y)

	hasil = 1/(3*x*x+10) + 10*y + 2

	fmt.Print("Hasilnya = ", hasil)

}
