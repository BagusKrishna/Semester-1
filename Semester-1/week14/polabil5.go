//program to make X pattern
package main

import "fmt"

func main() {
	var i, k, n, s, spasi int

	fmt.Scan(&n)
	//buat cek dia ganjil
	if n%2 == 1 {

		
		for i = 1; i <= n; i++ {
			if i == 1 || i == n {
				k = 2*n - 3
				fmt.Print(i)
				for k > 0 {
					fmt.Print(" ")
					k = k - 1

				}
				fmt.Print(i)
			} else {
				if i%2 == 0 {
					spasi = 2
					for s = 1; s <= spasi; s++ {
						fmt.Print(" ")
					}

					fmt.Print(i)
				} else {

					for s = 1; s <= 2; s++ {
						fmt.Print("  ")
					}
					fmt.Print(i)
				}
			}
			fmt.Println()
		}
	}
}
