package main

import "fmt"

func main() {
	var n, bil, sum int
	fmt.Scan(&n)
	if n != 0 {

		for i := 1; i <= n; i++ {
			fmt.Scan(&bil)
			sum = sum + bil

		}
		fmt.Println("Rata-rata = ", sum/n)
	}
}
