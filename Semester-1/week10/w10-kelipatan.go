package main

import "fmt"

func main() {
	var n int

	fmt.Println("Masukkan angkamu")
	fmt.Scan(&n)

	if n%15 == 0 {
		fmt.Print("Ini kelipatan 3 & 5")
	} else if n%3 == 0 {
		fmt.Print("Ini Kelipatan 3")

	} else if n%5 == 0 {
		fmt.Print("Ini Kelipatan 5")
	} else {
		fmt.Print("waduh, ini bukan kelipatan 3 dan 5 bro")
	}
}
