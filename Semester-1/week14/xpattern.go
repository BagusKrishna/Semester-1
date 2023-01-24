// program to make X pattern
package main

import "fmt"

func main() {
	var n, i, j int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if i == j || i+j == n+1 {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
