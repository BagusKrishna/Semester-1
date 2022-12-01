// program untuk mengetahui apakah tahun tersebut merupakan tahun kabisat atau tidak.
// syarat tahun kabisat adalah habis dibagi 4 namun tidak habis dibagai 100, atau tahun tersebut habis dibagi 400.
package main

import "fmt"

func main() {
	var year int
	var kabisat bool

	fmt.Println("tahun berapa?")
	fmt.Scan(&year)
	kabisat = (year%4 == 0 && year%100 != 0) || (year%400 == 0)
	fmt.Println("Apakah kabisat?? ", kabisat)
}
