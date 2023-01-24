//program untuk mengubah bilangan desimal ke biner
package main

import "fmt"

func main() {
	var bilanganDesimal int
	var biner string
	biner = ""
	fmt.Scan(&bilanganDesimal)
	for bilanganDesimal != 0 {
		if bilanganDesimal%2 == 1 {
			biner = "1" + biner

		} else {
			biner = "0" + biner

		}
		bilanganDesimal /= 2
	}
	fmt.Print(biner)
}