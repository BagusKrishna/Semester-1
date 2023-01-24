package main
import "fmt"
func main(){
	var x, i, j, k int
	fmt.Scan(&x)
	for j = 1; j<=x; j++{
		if j == 1 || j == x{
			for i = 1; i<=x; i++{
				fmt.Print(j, " ")
			}
		}else{
			fmt.Print(j)
			k = 2*x - 3
			for k > 0{
				fmt.Print(" ")
				k = k - 1
			}
			fmt.Print(j)
		}
		fmt.Println()
	}
}