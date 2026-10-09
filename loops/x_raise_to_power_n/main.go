package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)

	ans := 1
	for i := 1; i <= n; i++ {
		ans *= x
	}
	fmt.Println(ans)
}
