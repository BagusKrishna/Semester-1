// Program untuk mengetahui gol yang terbanyak dan gol yang tersedikit
package main

import "fmt"

func main() {
	var g1, g2, g3, g4 int

	fmt.Println("Berapa banyak gol hari ini? ")
	fmt.Scan(&g1, &g2, &g3, &g4)
	if g1 >= g2 && g1 >= g3 && g1 >= g4 {
		fmt.Print(g1)
	} else if g2 >= g1 && g2 >= g3 && g2 >= g4 {
		fmt.Print(g2)
	} else if g3 >= g1 && g3 >= g2 && g3 >= g4 {
		fmt.Print(g3)
	} else if g4 >= g1 && g4 >= g2 && g4 >= g3 {
		fmt.Print(g4)
	}
	if g1 <= g2 && g1 <= g3 && g1 <= g4 {
		fmt.Print(" ", g1)
	} else if g2 <= g1 && g2 <= g3 && g2 <= g4 {
		fmt.Print(" ", g2)
	} else if g3 <= g1 && g3 <= g2 && g3 <= g4 {
		fmt.Print(" ", g3)
	} else if g4 <= g1 && g4 <= g2 && g4 <= g3 {
		fmt.Print(" ", g4)
	}
}
